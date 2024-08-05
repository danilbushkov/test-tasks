<?php

namespace Tests\Feature;

// use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Http\UploadedFile;
use Illuminate\Support\Facades\Storage;
use Tests\TestCase;

class ProductTest extends TestCase
{
    use RefreshDatabase;

    public function test_import_products_without_file(): void {
        $response = $this->post('/api/products/import');
        $response->assertStatus(422)->assertJson([
            'error' => [
                'message' => 'Файл не найден'
            ],
        ]);
    }

    public function test_import_products_with_picture(): void {
        Storage::fake('avatars');

        $file = UploadedFile::fake()->image('avatar.jpg');
        $response = $this->post('/api/products/import', [
            'file' => $file
        ]);
        $response->assertStatus(422)->assertJson([
            'error' => [
                'message' => 'Файл не xlsx|xls формата'
            ]
        ]);
    }

    public function test_import_products_with_valid_xlsx_file(): void {
        $path = Storage::path('tests/import_example_small.xlsx');
        $file = new UploadedFile($path, 'import_example.xlsx', null, null, true);


        Storage::fake('local');
        $response = $this->post('/api/products/import', [
            'file' => $file
        ]);

        $response->assertStatus(204);

        $this->assertDatabaseCount('products', 6);
        $this->assertDatabaseCount('additional_product_fields', 100);
        $this->assertDatabaseCount('product_pictures', 20);

        $this->assertDatabaseHas('products',
            [
                'external_code' => '3UHfAid1jaMiwgBuNvnsf3',
                'price' => '1320.00',
                'description' => 'Мужские трикотажные домашние бермуды. Легкие, мягкие, дышащие и очень удобные. Универсальная длина ниже колена, пояс на эластичной ленте и карманы создают дополнительный комфорт.Отличный вариант для дома и отдыха. Бермуды стильно смотрятся благодаря эффектному рисунку и позволяют создать эффектный мужской образ.  Производство РОССИЯ',
                'name' => 'Бермуды мужские, Grigio/Verde, OMSA, 46(M), РОССИЯ'
            ]
        );
        $this->assertDatabaseHas('additional_product_fields', [
            'product_id' => 3,
            'key' => 'seo title',
            'value' => 'Купить Бандалетки жен., Daino, MINIMI, 6(XXL), КИТАЙ в интернет магазине creativeparadise.online',
        ]);

        $this->assertDatabaseHas('product_pictures', [
            'link' => 'http://catalog.collant.ru/pics/SNL-504038_b2.jpg',
            'product_id' => 1
        ]);

        Storage::disk('local')->assertExists('public/pictures/SNL-504038_m.jpg');
        Storage::disk('local')->assertExists('public/pictures/SNL-454043_b1.jpg');
        Storage::disk('local')->assertMissing('public/pictures/7_1.jpg');
    }
}
