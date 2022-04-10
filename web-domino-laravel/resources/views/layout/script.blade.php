<!-- jQuery 3 -->
<script src="/admin-lte/bower_components/jquery/dist/jquery.min.js"></script>
<!-- jQuery UI 1.11.4 -->
<script src="/admin-lte/bower_components/jquery-ui/jquery-ui.min.js"></script>
<!-- Resolve conflict in jQuery UI tooltip with Bootstrap tooltip -->
<script>
    $.widget.bridge('uibutton', $.ui.button);

    // csrf-token
    $.ajaxSetup({
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content'),
        }
    });
</script>
<!-- Bootstrap 3.3.7 -->
<script src="/admin-lte/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- Select2 -->
<script src="/admin-lte/bower_components/select2/dist/js/select2.full.min.js"></script>
<!-- Morris.js charts -->
<script src="/admin-lte/bower_components/raphael/raphael.min.js"></script>
<script src="/admin-lte/bower_components/morris.js/morris.min.js"></script>
<!-- Sparkline -->
<script src="/admin-lte/bower_components/jquery-sparkline/dist/jquery.sparkline.min.js"></script>
<!-- jvectormap -->
<script src="/admin-lte/plugins/jvectormap/jquery-jvectormap-1.2.2.min.js"></script>
<script src="/admin-lte/plugins/jvectormap/jquery-jvectormap-world-mill-en.js"></script>
<!-- jQuery Knob Chart -->
<script src="/admin-lte/bower_components/jquery-knob/dist/jquery.knob.min.js"></script>
<!-- daterangepicker -->
<script src="/admin-lte/bower_components/moment/min/moment.min.js"></script>
<script src="/admin-lte/bower_components/bootstrap-daterangepicker/daterangepicker.js"></script>
<!-- bootstrap time picker -->
<script src="/admin-lte/plugins/timepicker/bootstrap-timepicker.min.js"></script>
<!-- datepicker -->
<script src="/admin-lte/bower_components/bootstrap-datepicker/dist/js/bootstrap-datepicker.js"></script>
<!-- Bootstrap WYSIHTML5 -->
<script src="/admin-lte/plugins/bootstrap-wysihtml5/bootstrap3-wysihtml5.all.min.js"></script>
<!-- Slimscroll -->
<script src="/admin-lte/bower_components/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<!-- iCheck 1.0.1 -->
<script src="/admin-lte/plugins/iCheck/icheck.min.js"></script>
<!-- FastClick -->
<script src="/admin-lte/bower_components/fastclick/lib/fastclick.js"></script>
<!-- admin-lte App -->
<script src="/admin-lte/dist/js/admin-lte.min.js"></script>
<!-- DataTables -->
<script src="/admin-lte/bower_components/datatables.net/js/jquery.dataTables.min.js"></script>
<script src="/admin-lte/bower_components/datatables.net-bs/js/dataTables.bootstrap.min.js"></script>
<script>
    document.onload = function () {
        if ($('.select2').length > 0) $('.select2').select2();

        if (document.getElementById('modalSearchDateRangePicker')) {
            $('#modalSearchDateRangePicker').daterangepicker();
        }
    };
</script>
{{--<script src="/js/echarts/echarts.min.js"></script>--}}
<script src="/admin-lte/bower_components/ckeditor/ckeditor.js"></script>
{{--<script src="/js/tools.js"></script>--}}
<script type="text/javascript" src="/layer/layer.js"></script>
