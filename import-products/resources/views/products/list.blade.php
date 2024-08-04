@extends('layouts.app')

@section('title', 'Список товаров')

@section('content')
    <p>Список товаров</p>
    @foreach ($products as $product)
        <a href="/products/1">{{ $product->name }}</a><br/>
    @endforeach
@endsection
