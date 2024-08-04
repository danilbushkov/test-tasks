<?php

namespace App\Http\Controllers\Pages;

use App\Http\Controllers\Controller;
use App\Models\Product;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Illuminate\View\View;

class ProductController extends Controller
{
    public function list(Request $request): RedirectResponse | View {
        $page = $request->input('page');
        if($page !== null) {
            if(!ctype_digit($page) || (int)$page < 1) {
                return redirect('/products?page=1');
            }

        } else {
            $page = 1;
        }
        $maxItems = 10;
        $count = Product::count();
        $maxPage = ceil($count / $maxItems);
        $products = [];
        if($maxPage == 0 ) {
            if($page > 1) {
                return redirect('/products?page=1');
            }
        } else {
            if($page > $maxPage) {
                return redirect('/products?page='.$maxPage);
            }
            $products = Product::select('id', 'name')
                ->offset(($page-1)*$maxItems)
                ->limit($maxItems)->get();
        }



        return view('products.list', [
            'products' => $products,
            'page' => $page,
            'maxPage' => $maxPage
        ]);
    }

    public function import(): string {
        return view('products.import');
    }

    public function item(string $id): string {
        return view('products.item');
    }

}
