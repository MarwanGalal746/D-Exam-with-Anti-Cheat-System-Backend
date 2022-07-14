package dexam.desktopapplication.anticheat;

import dexam.desktopapplication.Main;
import dexam.desktopapplication.messaging.RabbitMq;

import java.io.IOException;
import java.util.ArrayList;
import java.util.concurrent.TimeUnit;

public class AntiCheat {
    public static void initiate(String browser) throws IOException, InterruptedException {

        ProcessManager processManager = new ProcessManager();
        ArrayList<String> p = new ArrayList<>();
        p.add("Discord.exe");
        //p.add("explorer.exe");
        p.add("idea64.exe");
        p.add("sihost.exe");
        p.add("cmd.exe");
        p.add("java.exe");
        p.add("javaw.exe");
        p.add("ApplicationFrameHost.exe");
        p.add("dwm.exe");
        processManager.addToIgnoreList(p);
        ArrayList<String> p2 = new ArrayList<>();
        p2.add("chrome.exe");
        p2.add("yarab.exe");
        p2.add("RabbitMqPublisher.exe");
        processManager.addToIgnoreList(p2);
        processManager.killProcesses();
        Thread newThread = new Thread(() -> {
            try {
                boolean result = processManager.monitorActivity();
                if(!result) {
                    RabbitMq.send("cheat-"+Main.userId);
                    String[] cmd = {"cmd.exe", "/c", "start explorer.exe"};
                    Runtime.getRuntime().exec(cmd);
                    System.exit(1);
                    Main.stopApp();
                }
            } catch (IOException e) {
                throw new RuntimeException(e);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        newThread.start();

        return;
    }
}
