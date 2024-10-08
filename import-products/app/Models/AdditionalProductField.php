<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class AdditionalProductField extends Model
{
    use HasFactory;

    protected $fillable = [
        'key',
        'value',
        'product_id'
    ];
}
