<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

use Illuminate\Support\Facades\Route;

Route::get("/", function () {
    return view("welcome");
});

// 权鉴
Route::prefix("authorization")->name("authorization:")->group(function () {
    Route::get("register", "AuthorizationController@GetRegister")->name("GetRegister");  // 注册页面
    Route::post("register", "AuthorizationController@PostRegister")->name("PostRegister");  // 注册
    Route::get("login", "AuthorizationController@GetLogin")->name("GetLogin");  // 登录页面
    Route::post("login", "AuthorizationController@PostLogin")->name("PostLogin");  // 登录
});
