@extends('layouts.app')

@section('title', 'Товар')

@section('content')
    <p>Товар</p>
    <div class = "product">
        <div class = "product-name">Название: {{ $product->name }}</div><br/>
        <div class = "product-description">
            Описание: {{ $product->description }}
        </div><br/>
        <div class = "product-external-code">
            Внешний код: {{ $product->external_code }}
        </div><br/>
        <div class = "product-price">
            Цена: {{ $product->price }}
        </div><br/>
        <div class = "product-discount">
            Скидка: {{ $product->discount }}
        </div><br/>
        <div class = "product-pictures">
            @foreach ($pictures as $picture)
                {{ $picture->path }}
            @endforeach
        </div><br/>
        <div class = "additional-product-field-title">
            Дополнительные поля:
        </div>
        <div class = "additional-product-fields">
            @foreach ($additionalFields as $field)
                <div class = "additional-product-field">
                    {{ $field->key }}: {{ $field->value }}
                </div><br/>
            @endforeach
        </div>
    </div>

@endsection
