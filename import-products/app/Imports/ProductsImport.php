<?php

namespace App\Imports;

use App\Models\Product;
use Maatwebsite\Excel\Concerns\PersistRelations;
use Maatwebsite\Excel\Concerns\ToModel;
use Maatwebsite\Excel\Concerns\WithHeadingRow;
use Maatwebsite\Excel\Imports\HeadingRowFormatter;



HeadingRowFormatter::default('none');

class ProductsImport implements ToModel, PersistRelations, WithHeadingRow
{
    /**
    * @param array $row
    *
    * @return \Illuminate\Database\Eloquent\Model|null
    */
    public function model(array $row): Product
    {
        $product = new Product([
            'external_code' => $row['Внешний код'],
            'name' => $row['Наименование'] ?? "",
            'description' => $row['Описание'] ?? "",
            'price' => str_replace(',', '.', $row['Цена: Цена продажи']),
            'discount' => array_key_exists('Скидка', $row) ? str_replace(',', '.', $row['Скидка']) : 0
        ]);

        return $product;
    }
}
