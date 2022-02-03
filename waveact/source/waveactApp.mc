import Toybox.Application;
import Toybox.Lang;
import Toybox.WatchUi;
import Toybox.Sensor;
import Toybox.System;
import Toybox.SensorLogging;
import Toybox.ActivityRecording;
import Toybox.Communications;
import Toybox.Graphics;
import Toybox.ActivityRecording;


class waveactApp extends Application.AppBase {

    var responseCode as Number?;

    function initialize() {
        AppBase.initialize();
        
    }

    function onStart(state as Dictionary?) as Void {
    }

    function onStop(state as Dictionary?) as Void {
    }

    function getInitialView() as Array<Views or InputDelegates>? {
        return [ new waveactView(), new waveactDelegate() ] as Array<Views or InputDelegates>;
    }

}

function getApp() as waveactApp {
    return Application.getApp() as waveactApp;
}