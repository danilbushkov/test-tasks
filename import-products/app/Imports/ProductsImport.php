<?php

namespace App\Imports;

use App\Models\AdditionalProductField;
use App\Models\Product;
use App\Models\ProductPicture;
use Illuminate\Support\Collection;
use Maatwebsite\Excel\Concerns\PersistRelations;
use Maatwebsite\Excel\Concerns\ToCollection;
use Maatwebsite\Excel\Concerns\ToModel;
use Maatwebsite\Excel\Concerns\WithHeadingRow;
use Maatwebsite\Excel\Imports\HeadingRowFormatter;



HeadingRowFormatter::default('none');

class ProductsImport implements ToCollection, WithHeadingRow
{

    public function collection(Collection $rows): Void
    {
        foreach ($rows as $row)
        {
            $product = Product::create([
                'external_code' => $row['Внешний код'],
                'name' => $row['Наименование'] ?? "",
                'description' => $row['Описание'] ?? "",
                'price' => str_replace(',', '.', $row['Цена: Цена продажи'] ?? 0),
                'discount' => str_replace(',', '.', $row['Скидка'] ?? 0)
            ]);

            foreach($row as $key => $value) {
                if(str_starts_with($key, 'Доп. поле:') && $value !== null) {
                    $key = str_replace("Доп. поле: ", "", $key);
                    if($key === 'Ссылки на фото') {
                        $links = explode(',', str_replace(' ', '', $value));
                        foreach($links as $link) {
                            ProductPicture::create([
                                'path' => "",
                                'link' => $link,
                                'product_id' => $product->id
                            ]);
                        }
                    } else {
                        AdditionalProductField::create([
                            'key' => $key,
                            'value' => $value,
                            'product_id' => $product->id
                        ]);
                    }
                }
            }
        }
    }
}

