@extends('layouts.app')

@section('title', 'Импортировать товары')

@section('content')
    <p>Импортировать товары</p>
    <input type="file" id="file-input">
    <br/>
    <button onClick="sendFile()">Загрузить</button>
    <div id="import-error"></div>


    <script src="{{ asset('api.js') }}"></script>
@endsection

