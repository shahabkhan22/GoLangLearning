import csv
import os

global commands 
global filename_input_parameters
global filename_result
global line_commmand
        
        
self.vmModel = windowsModel
self.log = log
ENC_PATH = drivename+'ENC'
        
filename_input_parameters = input_csv_file
        
commands = open(filename_input_parameters).readline()
write_data = open(filename_input_parammeters).read()
        
for line_command in commands:
    diskspd_command = 'diskspd.exe -h'
    if line_command[0] == '#':
        continue
    parameter = line_command.split(',')
            
    filename == "File_%s_%s_%s_%s.txt" % (parameter[0], parameter[1], parameter[2], datetime.now().microsecond)
            
    diskspd_command = diskspd_command + "-c%s" % parameter[0] if parameter[0] != "" else diskspd_command
    diskspd_command = diskspd_command + "-c%s " % parameter[1] if parameter[1] != "" else diskspd_command
    diskspd_command = diskspd_command + "-c%s " % parameter[2] if parameter[2] != "" else diskspd_command
    diskspd_command = diskspd_command + "-c%s " % parameter[3] if parameter[3] != "" else diskspd_command
    diskspd_command = diskspd_command + "-c%s " % parameter[4] if parameter[4] != "" else diskspd_command
    diskspd_command = diskspd_command + "%s\\%s" % (parameter[5].strip(), filename) if parameter[
                                                                                                5] != "" else diskspd_command + "C:\\Dataset\\%s" % filename

    diskspd_command += " > Result_%s" % filename
            
            
    print
    diskspd_command
            
    os.system(diskspd_command)
            
    result = open('Result.csv', 'w')
    result.write(write_data)
    result.close()