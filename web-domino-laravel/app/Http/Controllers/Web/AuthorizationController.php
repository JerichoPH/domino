<?php

namespace App\Http\Controllers\Web;

use App\Http\Controllers\Controller;
use App\Validations\Web\RegisterValidation;
use Curl\Curl;
use Illuminate\Contracts\View\Factory;
use Illuminate\Foundation\Application;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\View\View;
use function App\Factions\FailForbidden;
use function App\Factions\FailValidate;
use function App\Factions\OkDict;

class AuthorizationController extends Controller
{
    private $curl;
    private $apiRootUrl;

    public function __construct()
    {
        $this->curl = new Curl();
        $this->apiRootUrl = env("API_ROOT_URL");
    }

    /**
     * 登录页面
     * @return Factory|Application|View
     */
    public function GetLogin()
    {
        return view("authorization.login");
    }

    public function PostLogin(Request $request)
    {

    }

    /**
     * 注册页面
     * @return Factory|Application|View
     */
    public function GetRegister()
    {
        return view("authorization.register");
    }

    /**
     * 注册
     * @param Request $request
     * @return JsonResponse
     */
    public function PostRegister(Request $request): JsonResponse
    {
        $validation = new RegisterValidation($request);
        $v = $validation->check();
        if ($v->fails()) return FailValidate($v->errors()->first());

        $validated = $validation->validated();

        $this->curl->post("$this->apiRootUrl/v1/authorization/register", $validated->toArray());
        if ($this->curl->error) return FailForbidden($this->curl->response);

        return OkDict($this->curl->response);
    }
}
