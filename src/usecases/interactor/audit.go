package interactor

import (
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type AuditUseCase struct {
	AuditContract port.AuditContract
	AuditCrypt    port.AuditCrypt
	ContentRepo   port.ContentRepository
}

func NewAuditUseCase(auditContract port.AuditContract, auditCrypt port.AuditCrypt, contentRepo port.ContentRepository) *AuditUseCase {
	return &AuditUseCase{
		AuditContract: auditContract,
		AuditCrypt:    auditCrypt,
		ContentRepo:   contentRepo,
	}
}

func (au *AuditUseCase) ProofGen() (*entities.ProofList, error) {
	proofList := &entities.ProofList{}
	// * DBからコンテンツ情報を全取得
	receipts, err := au.ContentRepo.All()
	if err != nil {
		return nil, err
	}
	//* SPが所有する全てのコンテンツに対して証明データを生成する
	for i := 0; i < len(receipts); i++ {
		//* チャレンジをブロックチェーンから読み込む
		challen, err := au.AuditContract.GetChallen(receipts[i].ID)
		if err != nil {
			return nil, err
		}
		// //* チャレンジIDからハッシュデータをブロックチェーンから読み込む
		contentLog, err := au.AuditContract.Get(receipts[i].ID)
		if err != nil {
			return nil, err
		}
		//* proof作成
		proof, err := au.AuditCrypt.AuditProofGen(challen, receipts[i], contentLog) // TODO: 実装修正
		if err != nil {
			return nil, err
		}
		//* proofをブロックチェーンに登録
		err = au.AuditContract.RegisterProof(proof)
		if err != nil {
			return nil, err
		}

		proofList.DataList = append(proofList.DataList,
			entities.Proof{
				Myu:       proof.Myu,
				Gamma:     proof.Gamma,
				ContentId: proof.ContentId,
			},
		)
		proofList.Total += 1
	}
	return proofList, nil
}
