import Toybox.Graphics;
import Toybox.WatchUi;

class waveactView extends WatchUi.View {

    // CHANGEME
    public static const base_url = "https://xxxxx.com/";

    function initialize() {
        View.initialize();
    }

    function onLayout(dc as Dc) as Void {

        var options = {
            :period => 1,               // 1 second sample time
            :accelerometer => {
                :enabled => true,
                :sampleRate => 25       // 25 samples
            }
        };

        Sensor.registerSensorDataListener(method(:accel_callback), options);
        
        setLayout(Rez.Layouts.MainLayout(dc));
    }

    function onReceive(responseCode as Number, data as Dictionary?) {
        System.println(responseCode);
    }

    function makeRequest(sX as String, sY as String, sZ as String) as Void {
        var url = base_url + "data";

        var parameters = {
            "x" => sX,
            "y" => sY,
            "z" => sZ
        };

        var options = {
            :method => Communications.HTTP_REQUEST_METHOD_POST,
            :headers => {
            "Content-Type" => Communications.REQUEST_CONTENT_TYPE_JSON
            }
        };
        
        var responseCallback = method(:onReceive);
        Communications.makeWebRequest(url, parameters, options, method(:onReceive));
    }

    function healthcheck(){
        var url = base_url + "health";
        var parameters = null;
        var options = {
            :method => Communications.HTTP_REQUEST_METHOD_GET
        };
        var responseCallback = method(:onReceive);
        Communications.makeWebRequest(url, parameters, options, method(:onReceive));
    }

    function accel_callback(sensorData) {
        var x = sensorData.accelerometerData.x;
        var y = sensorData.accelerometerData.y;
        var z = sensorData.accelerometerData.z;
        
        makeRequest(x,y,z);
    }
    
    function onShow() as Void {
        healthcheck();
    }
    function onUpdate(dc as Dc) as Void {
        View.onUpdate(dc);
    }
    function onHide() as Void {
    }

}
