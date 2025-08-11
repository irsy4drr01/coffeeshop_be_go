package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedProductImages(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO product_images (product_id, slot_number, image_file)
		VALUES
			('4e841656-596c-434d-b5bb-1f27f5d7418c', '1', '1_1.webp'),
			('4e841656-596c-434d-b5bb-1f27f5d7418c', '2', '1_2.webp'),
			('4e841656-596c-434d-b5bb-1f27f5d7418c', '3', '1_3.webp'),
			('5740f5fe-5178-4933-9e7b-63cab72fd79a', '1', '2_1.webp'),
			('5740f5fe-5178-4933-9e7b-63cab72fd79a', '2', '2_2.webp'),
			('5740f5fe-5178-4933-9e7b-63cab72fd79a', '3', '2_3.webp'),
			('e6622f29-2a8b-4110-9a82-8616eed29570', '1', '3_1.webp'),
			('e6622f29-2a8b-4110-9a82-8616eed29570', '2', '3_2.webp'),
			('e6622f29-2a8b-4110-9a82-8616eed29570', '3', '3_3.webp'),
			('9a7b950f-3664-4df3-9da5-249e91de2b31', '1', '4_1.webp'),
			('9a7b950f-3664-4df3-9da5-249e91de2b31', '2', '4_2.webp'),
			('9a7b950f-3664-4df3-9da5-249e91de2b31', '3', '4_3.webp'),
			('8afc2e72-ef45-45b0-a936-62d34bd626bf', '1', '5_1.webp'),
			('8afc2e72-ef45-45b0-a936-62d34bd626bf', '2', '5_2.webp'),
			('8afc2e72-ef45-45b0-a936-62d34bd626bf', '3', '5_3.webp'),
			('7af264d9-fa31-45a4-8948-b2db4c267fd6', '1', '6_1.webp'),
			('7af264d9-fa31-45a4-8948-b2db4c267fd6', '2', '6_2.webp'),
			('7af264d9-fa31-45a4-8948-b2db4c267fd6', '3', '6_3.webp'),
			('f866f4f6-f89c-4395-90ca-241dfb52951c', '1', '7_1.webp'),
			('f866f4f6-f89c-4395-90ca-241dfb52951c', '2', '7_2.webp'),
			('f866f4f6-f89c-4395-90ca-241dfb52951c', '3', '7_3.webp'),
			('287ec09d-928c-4562-9f29-86ad95dce6f6', '1', '8_1.webp'),
			('287ec09d-928c-4562-9f29-86ad95dce6f6', '2', '8_2.webp'),
			('287ec09d-928c-4562-9f29-86ad95dce6f6', '3', '8_3.webp'),
			('95a70a1a-10c8-4a3f-80ad-6c430b74ef3e', '1', '9_1.webp'),
			('95a70a1a-10c8-4a3f-80ad-6c430b74ef3e', '2', '9_2.webp'),
			('95a70a1a-10c8-4a3f-80ad-6c430b74ef3e', '3', '9_3.webp'),
			('40425c10-f932-4b44-97a6-681b56a5ddfa', '1', '10_1.webp'),
			('40425c10-f932-4b44-97a6-681b56a5ddfa', '2', '10_2.webp'),
			('40425c10-f932-4b44-97a6-681b56a5ddfa', '3', '10_3.webp'),
			('0076aee4-9db2-4941-a69d-e07ff562dc3b', '1', '11_1.webp'),
			('0076aee4-9db2-4941-a69d-e07ff562dc3b', '2', '11_2.webp'),
			('0076aee4-9db2-4941-a69d-e07ff562dc3b', '3', '11_3.webp'),
			('aead3bb8-eaa3-4408-a192-cc36b227f464', '1', '12_1.webp'),
			('aead3bb8-eaa3-4408-a192-cc36b227f464', '2', '12_2.webp'),
			('aead3bb8-eaa3-4408-a192-cc36b227f464', '3', '12_3.webp'),
			('a2b74af4-06cc-4004-989e-6150af06926c', '1', '13_1.webp'),
			('a2b74af4-06cc-4004-989e-6150af06926c', '2', '13_2.webp'),
			('a2b74af4-06cc-4004-989e-6150af06926c', '3', '13_3.webp'),
			('dfd726cf-8c5c-44ac-8d06-82378bb4c31c', '1', '14_1.webp'),
			('dfd726cf-8c5c-44ac-8d06-82378bb4c31c', '2', '14_2.webp'),
			('dfd726cf-8c5c-44ac-8d06-82378bb4c31c', '3', '14_3.webp'),
			('31b9935a-7bd3-4ae0-898d-386e4cffb82e', '1', '15_1.webp'),
			('31b9935a-7bd3-4ae0-898d-386e4cffb82e', '2', '15_2.webp'),
			('31b9935a-7bd3-4ae0-898d-386e4cffb82e', '3', '15_3.webp'),
			('06bd407e-5fab-4590-9e3f-7dc442af3b42', '1', '16_1.webp'),
			('06bd407e-5fab-4590-9e3f-7dc442af3b42', '2', '16_2.webp'),
			('06bd407e-5fab-4590-9e3f-7dc442af3b42', '3', '16_3.webp'),
			('978feadd-1c68-479f-a99a-b831b732464b', '1', '17_1.webp'),
			('978feadd-1c68-479f-a99a-b831b732464b', '2', '17_2.webp'),
			('978feadd-1c68-479f-a99a-b831b732464b', '3', '17_3.webp'),
			('fad473ac-7dbe-470a-af07-836726e9b1a6', '1', '18_1.webp'),
			('fad473ac-7dbe-470a-af07-836726e9b1a6', '2', '18_2.webp'),
			('fad473ac-7dbe-470a-af07-836726e9b1a6', '3', '18_3.webp'),
			('e07cf18a-b204-455d-95e3-eaf489a805b2', '1', '19_1.webp'),
			('e07cf18a-b204-455d-95e3-eaf489a805b2', '2', '19_2.webp'),
			('e07cf18a-b204-455d-95e3-eaf489a805b2', '3', '19_3.webp'),
			('9c4817f0-455e-415f-9c79-8a7e6f4fc1ab', '1', '20_1.webp'),
			('9c4817f0-455e-415f-9c79-8a7e6f4fc1ab', '2', '20_2.webp'),
			('9c4817f0-455e-415f-9c79-8a7e6f4fc1ab', '3', '20_3.webp'),
			('40cb38bd-ba7f-4d40-b5e5-0a013b45e4c8', '1', '21_1.webp'),
			('40cb38bd-ba7f-4d40-b5e5-0a013b45e4c8', '2', '21_2.webp'),
			('40cb38bd-ba7f-4d40-b5e5-0a013b45e4c8', '3', '21_3.webp'),
			('97db8997-77ad-4793-8a83-e030ff84d4dd', '1', '22_1.webp'),
			('97db8997-77ad-4793-8a83-e030ff84d4dd', '2', '22_2.webp'),
			('97db8997-77ad-4793-8a83-e030ff84d4dd', '3', '22_3.webp'),
			('d6a142e4-9960-444f-9cca-041ceb595a9a', '1', '23_1.webp'),
			('d6a142e4-9960-444f-9cca-041ceb595a9a', '2', '23_2.webp'),
			('d6a142e4-9960-444f-9cca-041ceb595a9a', '3', '23_3.webp'),
			('9577276b-d4cb-4de2-a371-c4a25e790c63', '1', '24_1.webp'),
			('9577276b-d4cb-4de2-a371-c4a25e790c63', '2', '24_2.webp'),
			('9577276b-d4cb-4de2-a371-c4a25e790c63', '3', '24_3.webp'),
			('042875a5-da48-456a-a6e7-b6746f43ab02', '1', '25_1.webp'),
			('042875a5-da48-456a-a6e7-b6746f43ab02', '2', '25_2.webp'),
			('042875a5-da48-456a-a6e7-b6746f43ab02', '3', '25_3.webp'),
			('fb374b9a-ddde-4a0a-8739-325dcf9543dc', '1', '26_1.webp'),
			('fb374b9a-ddde-4a0a-8739-325dcf9543dc', '2', '26_2.webp'),
			('fb374b9a-ddde-4a0a-8739-325dcf9543dc', '3', '26_3.webp'),
			('ed6d3a40-7760-45cd-a8df-106daeef0227', '1', '27_1.webp'),
			('ed6d3a40-7760-45cd-a8df-106daeef0227', '2', '27_2.webp'),
			('ed6d3a40-7760-45cd-a8df-106daeef0227', '3', '27_3.webp'),
			('c504c3ea-af18-4bd8-a2e9-d7e773b1ea5d', '1', '28_1.webp'),
			('c504c3ea-af18-4bd8-a2e9-d7e773b1ea5d', '2', '28_2.webp'),
			('c504c3ea-af18-4bd8-a2e9-d7e773b1ea5d', '3', '28_3.webp'),
			('986f09ea-9843-45a3-91a2-da3193c12a63', '1', '29_1.webp'),
			('986f09ea-9843-45a3-91a2-da3193c12a63', '2', '29_2.webp'),
			('986f09ea-9843-45a3-91a2-da3193c12a63', '3', '29_3.webp'),
			('4c8c4d2e-6f9a-4c91-9f5d-0660ce5c3e6e', '1', '30_1.webp'),
			('4c8c4d2e-6f9a-4c91-9f5d-0660ce5c3e6e', '2', '30_2.webp'),
			('4c8c4d2e-6f9a-4c91-9f5d-0660ce5c3e6e', '3', '30_3.webp'),
			('8511cb06-1612-4b5a-9ef5-143abdc2077a', '1', '31_1.webp'),
			('8511cb06-1612-4b5a-9ef5-143abdc2077a', '2', '31_2.webp'),
			('8511cb06-1612-4b5a-9ef5-143abdc2077a', '3', '31_3.webp'),
			('40a2a975-9223-4a0e-9379-cb055fe8ae98', '1', '32_1.webp'),
			('40a2a975-9223-4a0e-9379-cb055fe8ae98', '2', '32_2.webp'),
			('40a2a975-9223-4a0e-9379-cb055fe8ae98', '3', '32_3.webp')
		ON CONFLICT DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed product_images: %v", err)
		return err
	}

	log.Println("Seeded product_images successfully.")
	return nil
}
