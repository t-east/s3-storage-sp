package interactor

import (
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type AuditHandler struct {
	AuditContract     port.AuditContract
	AuditCrypt        port.AuditCrypt
	ContentStorage    port.ContentStorage
	AuditRepository   port.AuditRepository
	ContentRepository port.ContentRepository
}

func NewAuditInputPort(
	contract port.AuditContract,
	crypt port.AuditCrypt,
	storage port.ContentStorage,
	repository port.AuditRepository,
	c_repo port.ContentRepository,
) port.AuditInputPort {
	return &AuditHandler{
		AuditContract:     contract,
		AuditCrypt:        crypt,
		ContentStorage:    storage,
		AuditRepository:   repository,
		ContentRepository: c_repo,
	}
}

func (ah *AuditHandler) ProofGen() (*entities.Proofs, error) {
	proofs := &entities.Proofs{}
	// * DBからコンテンツ情報を全取得
	receipts, err := ah.ContentRepository.All()
	if err != nil {
		return nil, err
	}
	//* SPが所有する全てのコンテンツに対して証明データを生成する
	for i := 0; i < len(receipts); i++ {
		//* チャレンジをブロックチェーンから読み込む
		challen, err := ah.AuditContract.GetChallen(receipts[i].ID)
		if err != nil {
			return nil, err
		}
		// //* チャレンジIDからハッシュデータをブロックチェーンから読み込む
		contentLog, err := ah.AuditContract.GetContentLog(receipts[i].ID)
		if err != nil {
			return nil, err
		}
		//* ファイルデータをストレージから読み込む
		content, err := ah.ContentStorage.Get(receipts[i].ID)
		if err != nil {
			return nil, err
		}
		//* proof作成
		proof, err := ah.AuditCrypt.AuditProofGen(challen, content, contentLog)
		if err != nil {
			return nil, err
		}
		//* proofをブロックチェーンに登録
		err = ah.AuditContract.RegisterProof(proof)
		if err != nil {
			return nil, err
		}
		//* proofをデータベースに登録
		updated, err := ah.AuditRepository.Update(proof)
		if err != nil {
			return nil, err
		}
		proofs.DataList = append(proofs.DataList,
			entities.Proof{
				Myu:       updated.Myu,
				Gamma:     updated.Gamma,
				ContentId: updated.ContentId,
			},
		)
		proofs.Total += 1
	}
	return proofs, nil
}
