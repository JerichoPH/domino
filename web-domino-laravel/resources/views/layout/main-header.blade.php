<header class="main-header">
    <!-- Logo -->
    <a href="{{url('/')}}" class="logo" style="background-color: #0477c4;">
        <!-- mini logo for sidebar mini 50x50 pixels -->
        <span class="logo-mini"><img src="/images/logo-mini.png" alt="" width="100%"></span>
        <!-- logo for regular state and mobile devices -->
        {{--<span class="logo-lg"><b style="font-size: 11px;">设备器材全生命周期管理系统</b><span style="font-size: 11px;">管理平台</span></span>--}}
        {{--        <span class="logo-lg"><img src="/images/logo{{ env('RAILWAY_CODE') ? '-'.env('RAILWAY_CODE') : '' }}.png" alt="" width="100%"></span>--}}
        <span class="logo-lg" style="text-align: left;font-size: 17px;color: #FFFFFF;">
            <img class="fa-image" style="width:35px;height: 35px;" src="/images/logo-mini.png">&nbsp;
            广州局集团公司
        </span>
    </a>
    <!-- Header Navbar: style can be found in header.less -->
    <nav class="navbar navbar-static-top" style="background-color: #0477c4;">
        <!-- Sidebar toggle button-->
        {{--<a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">--}}
        {{--    <span class="sr-only">Toggle navigation</span>--}}
        {{--</a>--}}
        <ul class="nav navbar-nav">
            <li class="dropdown tasks-menu">
                <a href="javascript:" data-toggle="push-menu" role="button" style="font-size: 18px;">
                    设备器材全生命周期管理系统
                    <small>{{ env('ORGANIZATION_NAME') }}</small>
                    {{--!env('IP_CONTROLLER') ? session('currentClientIp') : ''--}}
                </a>
            </li>
        </ul>

        <div class="navbar-custom-menu">
            <ul class="nav navbar-nav">
                <!-- Messages: style can be found in dropdown.less-->
                {{--<li class="dropdown tasks-menu">--}}
                {{--    <a href="javascript:" onclick="location.href='/warehouse/report/scanInBatch'" style="font-size: 22px;"><i class="fa fa-barcode"></i></a>--}}
                {{--</li>--}}
                {{--<li class="dropdown tasks-menu">--}}
                {{--    <a href="javascript:" onclick="fnModalSearch()" style="font-size: 22px;"><i class="fa fa-search">&nbsp;</i>老搜索</a>--}}
                {{--</li>--}}

                <li class="dropdown messages-menu" style="display: none;">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                        <i class="fa fa-envelope-o"></i>
                        <span class="label label-danger" id="liMessagesCount">0</span>
                    </a>
                    <ul class="dropdown-menu">
                        {{-- <li class="header">共<span id="spanMessagesCount">0</span>条消息</li> --}}
                        <li class="header">消息列表</li>
                        <li>
                            <!-- inner menu: contains the actual data -->
                            <ul class="menu" id="ulMessages"></ul>
                        </li>
                        {{-- <li class="footer"><a href="{{ url('message/input') }}">查看所有消息</a></li> --}}
                        {{--<li class="footer">共<span id="spanMessagesCount">0</span>条消息</li>--}}
                    </ul>
                </li>
                <li class="dropdown messages-menu">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                        <i class="fa fa-bell"></i>
                        <span class="label label-danger" id="spanScrapedAndCycleFixWarningCount">0</span>
                    </a>
                    <div class="dropdown-menu" style="width: 600px;">
                        <div id="divScrapedAndCycleFixWarnings"></div>
                    </div>
                    {{--<ul class="dropdown-menu" style="width: 1000px;">--}}
                    {{--    <li class="header">消息列表</li>--}}
                    {{--    <li><ul class="menu" id="ulScrapedAndCycleFixWarnings"></ul></li>--}}
                    {{--    <li class="footer">共<span id="spanScrapedAndCycleFixWarningCount">0</span>条消息</li>--}}
                    {{--</ul>--}}
                </li>
                <li class="dropdown tasks-menu">
                    <a href="{{ url('query') }}"><i class="fa fa-search">&nbsp;</i></a>
                </li>
                <li class="dropdown tasks-menu">
                    @switch(env('ORGANIZATION_CODE'))
                        @case('B041')
                        <a href="{{ url('newmonitor/') }}" target="_blank"><i class="fa fa-map-o">&nbsp;</i></a>
                        @break
                        @case('B048')
                        @case('B049')
                        @case('B050')
                        @case('B051')
                        @case('B052')
                        @case('B053')
                        @case('B074')
                        <a href="{{ url('monitor/') }}" target="_blank"><i class="fa fa-map-o">&nbsp;</i></a>
                        @break
                        @default
                        @break
                    @endswitch
                </li>
                <li class="dropdown messages-menu">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                        <i class="fa fa-television"></i>
                    </a>
                    @if(session('account.work_area_unique_code'))
                        <ul class="dropdown-menu">
                            {{-- <li class="header">共<span id="spanMessagesCount">0</span>条消息</li> --}}
                            <li class="header"><h4>出入所统计展板</h4></li>
                            <li>
                                <!-- inner menu: contains the actual data -->
                                <ul class="menu">
                                    <li><a href="{{ url('warehouseReportDisplayBoard',session('account.work_area_unique_code')) }}/showWarehouseReport/today" target="_blank">今日出入所统计</a></li>
                                    <li><a href="{{ url('warehouseReportDisplayBoard',session('account.work_area_unique_code')) }}/showWarehouseReport/week" target="_blank">本周出入所统计</a></li>
                                    <li><a href="{{ url('warehouseReportDisplayBoard',session('account.work_area_unique_code')) }}/showWarehouseReport/month" target="_blank">本月出入所统计</a></li>
                                    @switch(intval(substr(session('account.work_area_unique_code'),5)))
                                        @case(1)
                                        @case(2)
                                        <li><a href="{{ url('warehouseReportDisplayBoard',session('account.work_area_unique_code')) }}/showCycleFix" target="_blank">周期修计划年表</a></li>
                                        @break
                                    @endswitch
                                </ul>
                            </li>
                            {{--<li class="footer"><a href="{{ url('message/input') }}">查看所有消息</a></li> --}}
                            {{--<li class="footer">共<span id="spanMessagesCount">0</span>条消息</li>--}}
                        </ul>
                    @endif
                </li>

                {{--<li class="dropdown tasks-menu">--}}
                {{--    <a href="javascript:" onclick="fnModalScanQrCode()" style="font-size: 22px;"><i class="fa fa-qrcode"></i></a>--}}
                {{--</li>--}}
            <!-- User Account: style can be found in dropdown.less -->
                <li class="dropdown user user-menu">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                        <img src="/images/logo_sm_bg.png" class="user-image"
                             alt="{{ session('account.nickname') }}">
                        <span class="hidden-xs">{{ session('account.nickname') }}</span>
                    </a>
                    <ul class="dropdown-menu">
                        <!-- User image -->
                        <li class="user-header" style="background-color: #0477c4;">
                            {{--<img src="/images/account-avatar-lack.jpeg" onclick="location.href='/profile'"--}}
                            <img src="/images/logo_sm_bg.png"
                                 class="img-circle" alt="{{ session('account.nickname') }}">
                            <p>
                                {{ session('account.nickname') }} - 管理员
                                <small>{{ session('account.created_at') }}</small>
                            </p>
                        </li>
                        <!-- Menu Body -->
                    {{--<li class="user-body">--}}
                    {{--<div class="row">--}}
                    {{--<div class="col-xs-4 text-center">--}}
                    {{--<a href="#">Followers</a>--}}
                    {{--</div>--}}
                    {{--<div class="col-xs-4 text-center">--}}
                    {{--<a href="#">Sales</a>--}}
                    {{--</div>--}}
                    {{--<div class="col-xs-4 text-center">--}}
                    {{--<a href="#">Friends</a>--}}
                    {{--</div>--}}
                    {{--</div>--}}
                    {{--</li>--}}
                    <!-- Menu Footer-->
                        <li class="user-footer">
                            <div class="pull-left">
                                <a href="{{ url('/account') }}/{{ session('account.id') }}/edit" class="btn btn-default btn-flat">修改个人信息</a>
                            </div>
                            <div class="pull-right">
                                <a href="{{ url('/logout') }}" class="btn btn-default btn-flat">退出登录</a>
                            </div>
                        </li>
                    </ul>
                </li>
                <!-- Control Sidebar Toggle Button -->
                {{--<li>--}}
                {{--    <a href="#" data-toggle="control-sidebar"><i class="fa fa-gears"></i></a>--}}
                {{--</li>--}}
            </ul>
        </div>
    </nav>
</header>
