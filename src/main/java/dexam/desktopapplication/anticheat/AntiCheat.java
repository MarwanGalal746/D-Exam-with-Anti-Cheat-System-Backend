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
        p.add(browser);
        p.add("Discord.exe");
        p.add("cmd.exe");
        p.add("idea64.exe");
        p.add("sihost.exe");
        p.add("java.exe");
        p.add("javaw.exe");
        p.add("ApplicationFrameHost.exe");
        p.add("dwm.exe");
        p.add("com.docker.backend.exe");
        p.add("com.docker.extensions.exe");
        p.add("com.docker.dev-envs.exe");
        p.add("com.docker.vpnkit.exe");
        p.add("com.docker.cli.exe");
        p.add("com.docker.proxy.exe");
        p.add("docker.exe");
        p.add("com.docker.service");
        p.add("Docker Desktop.exe");
//        p2.add("chrome.exe");
        p.add("DesktopApplication.exe");
        p.add("wsl.exe");
        p.add("wslhost.exe");
        p.add("vmmem");
        p.add("vpnkit-bridge.exe");
        p.add("vmwp.exe");
        p.add("vmcompute.exe");
        p.add("node.exe");


        processManager.addToIgnoreList(p);

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
