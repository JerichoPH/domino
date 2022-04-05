<?php

namespace App\Console\Commands;

use Curl\Curl;
use Illuminate\Console\Command;

class InitDataCommand extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'data:init';
    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';
    private $curl;
    private $rootUrl = "http://127.0.0.1:8080";
    private $v1Url = "/v1";

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
        $this->curl = new Curl();
    }

    /**
     * Execute the console command.
     */
    final public function handle(): void
    {
        // 注册用户
        $this->curl->post("{$this->rootUrl}$this->v1Url/authorization/register", [
            "username" => "zhangsan",
            "nickname" => "张三",
            "password" => "123123",
            "password_check" => "123123",
        ]);
        $this->handleResponse("注册用户");

        // 登录
        $this->curl->post("{$this->rootUrl}$this->v1Url/authorization/login", [
            "username" => "zhangsan",
            "password" => "123123",
        ]);
        $this->handleResponse("登录");
        $jwt = $this->curl->response->content->token;
        dd("jwt:", $jwt);
    }

    /**
     * 处理请求响应
     * @param string $operationName
     */
    final private function handleResponse(string $operationName = "")
    {
        if ($this->curl->error) dd("{$operationName}错误：", $this->curl->getErrorMessage(), $this->curl->getErrorCode(), $this->curl->response);

        dump("{$operationName}成功：", $this->curl->response);
    }

    /**
     * @return Curl
     */
    public function GetCurl(): Curl
    {
        return $this->curl;
    }

    /**
     * @param Curl $curl
     * @return $this
     */
    final public function SetCurl(Curl $curl): self
    {
        $this->curl = $curl;
        return $this;
    }
}
