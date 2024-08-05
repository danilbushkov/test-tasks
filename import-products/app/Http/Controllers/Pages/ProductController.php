<?php

namespace App\Http\Controllers\Pages;

use App\Http\Controllers\Controller;
use App\Models\Product;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Illuminate\View\View;

class ProductController extends Controller
{
    public function list(Request $request): View|RedirectResponse {
        $page = $request->input('page');
        if($page !== null) {
            if(!ctype_digit($page)) {
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

    public function import(): View {
        return view('products.import');
    }

    public function item(string $id): View|RedirectResponse {
        if(!ctype_digit($id)) {
            return redirect('/products?page=1');
        }

        $product = Product::find($id);
        $pictures = $product->pictures;
        $additionalFields = $product->additionalFields;

        if($product === null) {
            return redirect('/products?page=1');
        }
        return view('products.item', [
            'product' => $product,
            'pictures' => $pictures,
            'additionalFields' => $additionalFields,
        ]);
    }

}
