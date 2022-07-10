package dexam.desktopapplication.utility;

import java.util.Map;

public class BrowsersList {
    private static final Map<String, String> browsers = Map.ofEntries(
            Map.entry("Chrome", "chrome.exe"),
            Map.entry("Firefox", "firefox.exe"),
            Map.entry("Opera", "opera.exe"),
            Map.entry("Edge", "msedge.exe")
    );

    public static Map<String, String> getBrowsersList() {
        return browsers;
    }
}
