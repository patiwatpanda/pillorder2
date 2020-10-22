
import React, { FC, useEffect, useState } from 'react';
import TextField from '@material-ui/core/TextField';
//import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
//import Checkbox from '@material-ui/core/Checkbox';
//import CheckBoxOutlineBlankIcon from '@material-ui/icons/CheckBoxOutlineBlank';
//import CheckBoxIcon from '@material-ui/icons/CheckBox';
import Swal from 'sweetalert2'; // alert
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
//import moment from 'moment';
//import Chip from '@material-ui/core/Chip';
import {
  Container,
} from '@material-ui/core';
import 'date-fns';
//import DateFnsUtils from '@date-io/date-fns';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';


import { DefaultApi } from '../../services/api/apis'; // Api Gennerate From Command
import { EntEmployee } from '../../services/api/models/EntEmployee'; // import interface Employee
import { EntDentist } from '../../services/api/models/EntDentist'; // import interface Dentist
import { EntPatient } from '../../services/api/models/EntPatient'; // import interface Patient
import { EntPillorderItem } from '../../services/api/models/EntPillorderItem'; // import interface PillorderItem
//import { Container } from '@material-ui/core';
import { 
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
 
} from '@backstage/core';


const useStyles = makeStyles((theme) => ({
 
  button: {
    margin: theme.spacing(1, 0),
  },
   formControl: {
    width: 300,
  },

  noLabel: {
    marginTop: theme.spacing(3),
  },
  margin: {
    margin: theme.spacing(3),
  },
  textField: {
    width: '20ch',
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
}));

interface welcomePage {
 employee: number;
 dentist: number ;
 patient: number  ;
 pillordernameid: string;
 pillorderitem: number ;
 pillorderdate: string ;
  // create_by: number;
}


const WelcomePage: FC<{}> = () => {

  const classes = useStyles();

  const http = new DefaultApi();
  const [pillorder, setPillorder] = React.useState<
  Partial<welcomePage>
  >({});

  const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
  const [dentists, setDentists] = React.useState<EntDentist[]>([]);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const [pillorderitems, setPillorderItems] = React.useState<EntPillorderItem[]>([]);
  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  const getEmployees = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployees(res);
  };
  const getDentists = async () => {
    const res = await http.listDentist({ limit: 10, offset: 0 });
    setDentists(res);
  };
  const getPatients = async () => {
    const res = await http.listPatient({ limit: 10, offset: 0 });
    setPatients(res);
  };
  const getPilllorderItems = async () => {
    const res = await http.listPillorderitem({ limit: 10, offset: 0 });
    setPillorderItems(res);
  };
  useEffect(() => {
    getEmployees();
    getDentists();
    getPatients();
    getPilllorderItems();
  }, []);
// set data to object pillorder
    const handleChange = (
      event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
      const name = event.target.name as keyof typeof WelcomePage;
      const { value } = event.target;
      setPillorder({ ...pillorder, [name]: value });
      
      console.log(pillorder);
    };
    // clear input form
    function clear() {
      setPillorder({});
    }
     // function save data
    function save() {
      pillorder.pillorderdate += ":00+07:00";
      const apiUrl = 'http://localhost:8080/api/v1/pillorders';
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(pillorder),
      };
    
      console.log(pillorder); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
  
      fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data != undefined) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
    }
  

 const profile = { givenName: 'Detal System' };
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`Welcome ${profile.givenName || 'to Backstage'}`}
      subtitle="ระบบทันตกรรมของคลินิกทันตกรรม"  // <dd> <ComponanceTable></ComponanceTable></dd>
     ></Header>
     <Content>
  
     <div style={{textAlign: "center"}}>
    
       <ContentHeader title="ระบบสั่งยา" >
        
       </ContentHeader>
      <br></br>
      <Grid item xs={9}>
      <FormControl variant="outlined" className={classes.formControl}>
               <TextField 
               label="หมายเลขใบสั่งยา" 
               variant="outlined" 
               name="pillordernameid"
               type="string"
               value={pillorder.pillordernameid || ''}
               onChange={handleChange}
               />
             </FormControl>
           
     <br></br>
      <br></br>
      </Grid>
      <Grid item xs={9}>
    <FormControl variant="outlined" className={classes.formControl}>
        <InputLabel >รหัสผู้ป่วย</InputLabel>
        <Select
          name="patient"
          value={pillorder.patient || ''} // (undefined || '') = ''
          onChange={handleChange}
          
        >
           {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.patientName}
                      </MenuItem>
                    );
            })}

          
        </Select>
        </FormControl>
      </Grid>&emsp; &emsp;
        <Grid item xs={9}>
    <FormControl variant="outlined" className={classes.formControl}>
        <InputLabel >ยา</InputLabel>
        <Select
          name="pillorderitem"
          value={pillorder.pillorderitem || ''} // (undefined || '') = ''
          onChange={handleChange}
          
        >
           {pillorderitems.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.pillorderItemName}
                      </MenuItem>
                    );
            })}
            </Select>
      </FormControl>
      </Grid>
      <br></br>
    <br></br>
    <Grid item xs={9}>
    <FormControl variant="outlined" className={classes.formControl}>
        <InputLabel >ทันตแพทย์ผู้สั่งยา</InputLabel>
        <Select
          name="dentist"
          value={pillorder.dentist || ''} // (undefined || '') = ''
          onChange={handleChange}
        >
           {dentists.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.dentistName}
                      </MenuItem>
                    );
            })}          
        </Select>
      </FormControl>
      </Grid>      
      <br></br>
       
      <br></br>
      <Grid item xs={9}>
      <FormControl variant="outlined" className={classes.formControl}>
        <InputLabel >พนักงานผู้บันทึกการสั่งยา</InputLabel>
        <Select
          name="employee"
          value={pillorder.employee || ''} // (undefined || '') = ''
          onChange={handleChange}
        >
         {employees.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.employeeName}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </Grid>
      <br></br>
      <Grid item xs={9}>
      <FormControl variant="outlined" className={classes.formControl}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันที่"
                  name="pillorderdate"
                  type="datetime-local"
                  value={pillorder.pillorderdate|| ''} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
                
      </FormControl>
            </Grid>
      <br></br>

      <br></br>
      <Grid item xs={3}></Grid>
            <Grid item xs={9}>
      <Button
            variant="contained"
            color="primary"
            size="large"
            startIcon={<SaveIcon />}
            onClick={save}
      >
        สั่งยา
      </Button>&emsp;
      
      <Link component={RouterLink} to="/show_pillorder">
              <Button
                variant="contained"
                color="default"
              >
                แสดงข้อมูล
              </Button>
              </Link>
      </Grid>                
     </div>
     </Content>
  </Page>
 );
};
export default WelcomePage;
      