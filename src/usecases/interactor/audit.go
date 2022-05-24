package interactor

import (
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type AuditHandler struct {
	AuditContract   port.AuditContract
	AuditCrypt      port.AuditCrypt
	ContentStorage  port.ContentStorage
	AuditRepository port.AuditRepository
}

func NewAuditInputPort(
	contract port.AuditContract,
	crypt port.AuditCrypt,
	storage port.ContentStorage,
	repository port.AuditRepository,
) port.AuditInputPort {
	return &AuditHandler{
		AuditContract:   contract,
		AuditCrypt:      crypt,
		ContentStorage:  storage,
		AuditRepository: repository,
	}
}

func (ah *AuditHandler) Challen() (*entities.Proofs, error) {
	proofs := &entities.Proofs{}
	//* ブロックチェーンからcontentIDを全て読み込む
	var ids []string = []string{"a", "a", "a", "a", "a"}
	for i := 0; i < len(ids); i++ {
		//* チャレンジをブロックチェーンから読み込む
		challen, err := ah.AuditContract.GetChallen(ids[i])
		if err != nil {
			return nil, err
		}
		// //* チャレンジIDからハッシュデータをブロックチェーンから読み込む
		contentLog, err := ah.AuditContract.GetContentLog(ids[i])
		if err != nil {
			return nil, err
		}
		//* ファイルデータをストレージから読み込む
		content, err := ah.ContentStorage.Get(ids[i])
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
