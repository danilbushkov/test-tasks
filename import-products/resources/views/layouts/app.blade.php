<!DOCTYPE html>
<html>
    <head>
        <title>@yield('title')</title>
    </head>
    <body>
        <header>
        <a href="/products/">Список</a>
        <a href="/products/import">Импортировать</a>
        </header>

        <div class="container">
            @yield('content')
        </div>
    </body>
</html>
