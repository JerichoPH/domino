<?php

namespace App\Console\Commands;

use Curl\Curl;
use Illuminate\Console\Command;

class TestCommand extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'test';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     */
    final public function handle(): void
    {
        $client = new Curl();
        // $client->post("http://127.0.0.1:8080/v1/authorization/register", [
        //     "username" => "zhangsan",
        //     "password" => "123123",
        //     "password_check" => "123123",
        //     "nickname" => "张三",
        // ]);
        // if ($client->error) dd("错误", $client->getErrorMessage(), $client->getErrorCode(), $client->response);

        $client->post("http://127.0.0.1:8080/v1/authorization/login", [
            "username" => "zhangsan",
            "password" => "123123",
        ]);

        dd($client->response);
    }
}
