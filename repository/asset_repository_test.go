package repository_test

// func TestAssetRepository_CreateAsset_When_Success(t *testing.T) {
// 	ctx := context.Background()
// 	var assetExpected domain.Asset
// 	id, err := uuid.Parse("642fc397-abec-4e1e-8473-69803dbb9016")
// 	duration, err := time.Parse("01/01/0001", "01/01/0001")
// 	spec := []byte(`{"ram":"4GB","brand":"acer"}`)

// 	dummy := domain.Asset{
// 		Id:             id,
// 		Status:         "active",
// 		Category:       "laptop",
// 		PurchaseAt:     duration,
// 		PurchaseCost:   45000.00,
// 		Name:           "aspire-5",
// 		Specifications: spec,
// 	}

// 	config.Init()
// 	repository.InitDB()
// 	db := repository.GetDB()

// 	tx := db.MustBegin()
// 	tx.MustExec("delete from assets")
// 	tx.Commit()

// 	assetRepo := repository.NewAssetRepository()

// 	asset, err := assetRepo.CreateAsset(ctx, &dummy)

// 	fmt.Println()
// 	db.Get(&assetExpected, "SELECT * FROM assets WHERE id = $1", id)
// 	fmt.Println(assetExpected)

// 	assert.Equal(t, &assetExpected, asset)
// 	assert.Nil(t, err)

// 	fmt.Println()
// }

// func TestAssetRepository_CreateAsset_When_ReturnsError(t *testing.T) {
// 	ctx := context.Background()
// 	var assetExpected domain.Asset
// 	id, err := uuid.Parse("642fc397-abec-4e1e-8473-69803dbb9016")
// 	duration, err := time.Parse("01/01/0001", "01/01/0001")
// 	spec := []byte(`{"ram":"4GB","brand":"acer"}`)

// 	dummy := domain.Asset{
// 		Id:             id,
// 		Status:         "active",
// 		Category:       "laptop",
// 		PurchaseAt:     duration,
// 		PurchaseCost:   45000.00,
// 		Name:           "aspire-5",
// 		Specifications: spec,
// 	}

// 	config.Init()
// 	repository.InitDB()
// 	db := repository.GetDB()

// 	tx := db.MustBegin()
// 	tx.MustExec("delete from assets")
// 	tx.Commit()

// 	assetRepo := repository.NewAssetRepository()

// 	asset, err := assetRepo.CreateAsset(ctx, &dummy)

// 	fmt.Println()
// 	db.Get(&assetExpected, "SELECT * FROM assets WHERE id = $1", id)
// 	fmt.Println(assetExpected)

// 	assert.Equal(t, &assetExpected, asset)
// 	assert.Nil(t, err)
// 	fmt.Println()
// }

// func TestAssetRepository_GetAsset_When_Success(t *testing.T) {
// 	ctx := context.Background()
// 	var assetExpected domain.Asset
// 	id, err := uuid.Parse("642fc397-abec-4e1e-8473-69803dbb9016")
// 	if err != nil {
// 		fmt.Printf("Error while parsing uuid: %s", err.Error())
// 		return
// 	}
// 	config.Init()
// 	repository.InitDB()
// 	db := repository.GetDB()

// 	assetRepo := repository.NewAssetRepository()

// 	asset, err := assetRepo.GetAsset(ctx, id)

// 	fmt.Println()
// 	db.Get(&assetExpected, "SELECT * FROM assets WHERE id = $1", id)
// 	fmt.Println(assetExpected)

// 	assert.NotNil(t, asset)
// 	assert.Equal(t, &assetExpected, asset)
// 	assert.Nil(t, err)
// 	fmt.Println()
// }

// func TestAssetRepository_GetAsset_When_ReturnsError(t *testing.T) {
// 	ctx := context.Background()
// 	id, err := uuid.Parse("642fc397-abec-4e1e-8473-03dbb9017")
// 	if err != nil {
// 		fmt.Printf("Error while parsing uuid: %s", err.Error())
// 		return
// 	}

// 	assetRepo := repository.NewAssetRepository()

// 	asset, err := assetRepo.GetAsset(ctx, id)

// 	fmt.Println()
// 	assert.Nil(t, asset)
// 	assert.NotNil(t, err)
// 	fmt.Println()
// }

// func TestAssetRepository_GetAsset_When_AssetDoesNotExist(t *testing.T) {
// 	ctx := context.Background()
// 	id, err := uuid.Parse("642fc397-abec-4e1e-8473-69803dbb9017")
// 	if err != nil {
// 		fmt.Printf("Error while parsing uuid: %s", err.Error())
// 		return
// 	}

// 	assetRepo := repository.NewAssetRepository()

// 	asset, err := assetRepo.GetAsset(ctx, id)

// 	assert.Nil(t, asset)
// 	assert.NotNil(t, err)
// 	fmt.Println()
// }
