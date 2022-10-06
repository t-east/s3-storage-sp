package interactor

import (
	entities "sp/src/domains/entities"
	port "sp/src/usecases/port"
)

type AuditUseCase struct {
	auditContract   port.AuditContractPort
	contentContract port.ContentContractPort
	cryptPort       port.CryptPort
	ContentRepo     port.ContentRepository
}

func NewAuditUseCase(auditContract port.AuditContractPort, contentContract port.ContentContractPort, cryptPort port.CryptPort, contentRepo port.ContentRepository) *AuditUseCase {
	return &AuditUseCase{
		auditContract:   auditContract,
		contentContract: contentContract,
		cryptPort:       cryptPort,
		ContentRepo:     contentRepo,
	}
}

func (au *AuditUseCase) ProofGen() (*entities.ProofList, error) {
	proofList := &entities.ProofList{}
	// * DBからコンテンツ情報を全取得
	receipts, err := au.ContentRepo.List()
	if err != nil {
		return nil, err
	}
	//* SPが所有する全てのコンテンツに対して証明データを生成する
	for i := 0; i < len(receipts); i++ {
		//* チャレンジをブロックチェーンから読み込む
		challen, err := au.auditContract.GetChallen(receipts[i].ID)
		if err != nil {
			return nil, err
		}
		// //* チャレンジIDからハッシュデータをブロックチェーンから読み込む
		contentLog, err := au.contentContract.FindByID(receipts[i].ID)
		if err != nil {
			return nil, err
		}
		//* proof作成
		proof, err := au.cryptPort.AuditProofGen(challen, receipts[i], contentLog) // TODO: 実装修正
		if err != nil {
			return nil, err
		}
		//* proofをブロックチェーンに登録
		err = au.auditContract.RegisterProof(proof)
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
