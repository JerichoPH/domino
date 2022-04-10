<?php

namespace App\Exceptions;

use App\Facades\JsonResponseFacade;
use Exception;
use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Foundation\Application;
use Illuminate\Foundation\Exceptions\Handler as ExceptionHandler;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Routing\Redirector;
use Illuminate\Support\Facades\Log;
use Symfony\Component\HttpFoundation\Response as HttpFoundationResponse;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;
use function App\Factions\FailEmpty;
use function App\Factions\FailForbidden;
use function App\Factions\FailUnLogin;
use function App\Factions\FailUnAuthorization;
use function App\Factions\FailValidate;

class Handler extends ExceptionHandler
{
    /**
     * A list of the exception types that are not reported.
     *
     * @var array
     */
    protected $dontReport = [
        //
    ];
    /**
     * A list of the inputs that are never flashed for validation exceptions.
     *
     * @var array
     */
    protected $dontFlash = [
        'password',
        'password_confirmation',
    ];
    private $code = 0;

    /**
     * Report or log an exception.
     *
     * @param Exception $e
     * @return void
     */
    public function Report(Exception $e)
    {
        $this->code = time() . str_pad(rand(0, 9999), 2, '0', 0);

        Log::error($e->getMessage(), [
            'request' => request()->all(),
            'code' => $this->code,
            'trace' => $e->getTraceAsString(),
        ]);
    }

    /**
     * Render an exception into an HTTP response.
     * @param Request $request
     * @param Exception $e
     * @return Application|JsonResponse|RedirectResponse|Response|Redirector|HttpFoundationResponse
     */
    public function render($request, Exception $e)
    {
        $e_msg = $e->getMessage();
        $msg = "错误：{$e_msg}。错误代码：$this->code";

        if (env('APP_DEBUG') && $request->ajax()) dd($e);

        if ($e instanceof UnAuthorizationException) {
            return $request->ajax()
                ? FailUnAuthorization($msg)
                : back()->withInput()->with('danger', $msg);
        }

        if ($e instanceof EmptyException) {
            return $request->ajax()
                ? FailEmpty($msg)
                : back()->withInput()->with('danger', $msg);
        }

        if ($e instanceof ForbiddenException) {
            return $request->ajax()
                ? FailForbidden($msg)
                : back()->withInput()->with('danger', $msg);
        }

        if ($e instanceof UnLoginException) {
            return $request->ajax()
                ? FailUnLogin($msg)
                : redirect('/login', $msg);
        }

        if ($e instanceof UnOwnerException) {
            return $request->ajax()
                ? FailUnAuthorization($msg)
                : back()->withInput()->with('danger', $msg);
        }

        if ($e instanceof ValidateException) {
            return $request->ajax()
                ? FailValidate($msg)
                : back()->withInput()->with('danger', $msg);
        }

        if ($e instanceof ModelNotFoundException) {
            $msg = "错误：资源不存在。错误代码：$this->code";
            return $request->ajax()
                ? FailEmpty($msg)
                : back()->withInput()->with("danger", $msg);
        }

        if ($e instanceof NotFoundHttpException) {
            $msg = "错误：路由不存在。错误代码：$this->code";
            return $request->ajax()
                ? FailEmpty($msg)
                : back()->withInput()->with("danger", $msg);
        }

        if ($e instanceof ExcelInException) {
            return $request->ajax()
                ? FailForbidden($msg)
                : back()->with("danger", $msg);
        }

        if ($e instanceof Exception) {
            return $request->ajax()
                ? response()->json([
                    "msg" => "意外错误。错误代码：$this->code",
                    "status" => 500,
                    "errorCode" => 500,
                    "details" => [
                        "exception_type" => get_class($e),
                        "message" => $e->getMessage(),
                        "file" => $e->getFile(),
                        "line" => $e->getLine(),
                    ],
                ])
                : back()->withInput()->with("danger", "意外错误。错误代码：$this->code");
        }

        return parent::render($request, $e);
    }
}
