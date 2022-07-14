package dexam.desktopapplication.anticheat;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.concurrent.TimeUnit;

public class ProcessManager {
    private String processLineInfo;
    private String[] cmd = {"cmd.exe", "/c", "tasklist /FO csv | sort"};
    private BufferedReader processList;
    private ArrayList<String> ignoreList = new ArrayList<>();
    private ArrayList<String> afterKill = new ArrayList<>();

    private BufferedReader getProcessList() throws IOException {
        Process p = Runtime.getRuntime().exec(cmd);
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(p.getInputStream()));
        return bufferedReader;
    }

    private String getExectuableProcess(BufferedReader processList) throws IOException {
        while (!processLineInfo.contains(".exe") && !processLineInfo.contains(".EXE"))
            processLineInfo = processList.readLine();
        return processLineInfo;
    }

    public void killProcesses() throws IOException {
        processList = getProcessList();
        processLineInfo = processList.readLine();
        while (true) {
            processLineInfo = getExectuableProcess(processList);
            String[] arr = processLineInfo.split(",");
            String programName = arr[0].replaceAll("\"", "");
            ArrayList<String> singleProgramProcesses = getSingleProgramProcesses(programName);
            executeTaskKill(singleProgramProcesses);
            singleProgramProcesses.clear();
            if (processLineInfo == null)
                break;
        }
    }

    private ArrayList<String> getSingleProgramProcesses(String programName) throws IOException {
        boolean isService = false;
        ArrayList<String> singleProgramProcesses = new ArrayList<>();
        while (processLineInfo.contains(programName)) {
            singleProgramProcesses.add(programName);
            if (processLineInfo.contains("Services")) {
                isService = true;
            }
            processLineInfo = processList.readLine();
            if (processLineInfo == null)
                return singleProgramProcesses;
        }
        if (isService) {
            singleProgramProcesses.clear();
        }
        return singleProgramProcesses;
    }

    private void executeTaskKill(ArrayList<String> singleProgramProcesses) throws IOException {
        for (int i = 0; i < singleProgramProcesses.size(); i++) {
            String process = singleProgramProcesses.get(i);
            Runtime rt = Runtime.getRuntime();
            if (!checkIgnoreList(process)) {
                rt.exec("taskkill /f /im " + process);
            } else {
                break;
            }
        }
    }

    private boolean checkIgnoreList(String processName) {
        return ignoreList.contains(processName);
    }

    public void addToIgnoreList(ArrayList<String> programNames) {
        ignoreList.addAll(programNames);
    }

    public void clearIgnoreList() {
        ignoreList.clear();
    }

    public void removeFromIgnoreListByName(String programName) {
        ignoreList.remove(new String(programName));
    }

    public boolean monitorActivity() throws IOException, InterruptedException {
        TimeUnit.SECONDS.sleep(1);
        while (true) {
            boolean cond = false;
            ArrayList<String> result = new ArrayList<>();
            processList = getProcessList();
            processLineInfo = processList.readLine();
            while (true) {
                processLineInfo = getExectuableProcess(processList);
                String[] arr = processLineInfo.split(",");
                String programName = arr[0].replaceAll("\"", "");
                System.out.println(arr[0]);
                if (programName.equals("kill.exe")) {
                    cond = true;
                }
                processLineInfo = processList.readLine();
                if (processLineInfo == null)
                    break;
            }
            if (!cond) {
                return false;
            }
        }
    }
}