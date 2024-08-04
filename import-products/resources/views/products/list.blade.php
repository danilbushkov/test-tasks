@extends('layouts.app')

@section('title', 'Список товаров')

@section('content')
    <p>Список товаров</p>
    @foreach ($products as $product)
        <a href="/products/1">{{ $product->id }} {{ $product->name }}</a><br/>
    @endforeach
    @if ($page > 1)
        <a href="/products?page={{$page-1}}">{{"<"}}</a>
    @endif
    @if ($maxPage != 0 && $page != 1)
        <a href="/products?page=1">{{ 1 }}</a>
    @endif
    @if ($maxPage != 0)
        <span>{{ $page }}</span>
    @endif
    @if ($maxPage != 0 && $page != $maxPage)
        <a href="/products?page={{$maxPage}}">{{ $maxPage }}</a>
    @endif
    @if ($page < $maxPage)
        <a href="/products?page={{$page+1}}">{{">"}}</a>
    @endif
@endsection
