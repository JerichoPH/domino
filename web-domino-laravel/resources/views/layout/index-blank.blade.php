<!DOCTYPE html>
<html>
<head>
    @include('layout.head')
    <meta name="csrf-token" content="{{ csrf_token() }}">
    @yield('style')
</head>
<body class="hold-transition skin-blue-light sidebar-mini">
<div class="wrapper">
    @yield('content')
</div>
@include('layout.script')
@yield('script')
</body>
</html>
