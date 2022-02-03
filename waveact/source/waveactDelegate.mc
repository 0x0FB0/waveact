import Toybox.Lang;
import Toybox.WatchUi;
import Toybox.Communications;

class waveactDelegate extends WatchUi.BehaviorDelegate {

    function initialize() {
        BehaviorDelegate.initialize();
    }

    function onMenu() as Boolean {
        return true;
    }

    public function onSelect() as Boolean {
        return true;
    }

}