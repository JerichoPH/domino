<?php

namespace App\Validations\Web;

use App\Validations\Validation;
use Illuminate\Http\Request;

class RegisterValidation extends Validation
{
    public $rules = [
        "username" => ["required", "between:1,64",],
        "nickname" => ["required", "between:1,64",],
        "password" => ["required", "between:6,64",],
        "password_confirmation" => ["required",],
    ];

    public $messages = [];

    public $attributes = [
        "username" => "账号",
        "nickname" => "昵称",
        "password" => "密码",
        "password_confirmation"=>"确认密码",
    ];

    public function __construct(Request $request)
    {
        parent::__construct($request);

        $this->messages = array_merge(parent::$BASE_MESSAGES, $this->messages);
    }
}
