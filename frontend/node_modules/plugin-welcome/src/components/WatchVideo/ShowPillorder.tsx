
import React, { FC, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../services/api/apis';
import { EntPillorder } from '../../services/api/models'; // import interface Pillorder
import Button from '@material-ui/core/Button';

import moment from 'moment';
import { 
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
  
 } from '@backstage/core';
// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
  table: {
    minWidth: 650,
  },
}));



const ShowPillorder: FC<{}> = () => {
  
  
  const classes = useStyles();
  const http = new DefaultApi();
  const [pillorders, setPillorders] = useState<EntPillorder[]>([]);;
  const [loading, setLoading] = useState(true);

// get pillorders
useEffect(() => {
  const getPillorders = async () => {
    const res = await http.listPillorder({ limit: 20, offset: 0 });
    setLoading(true);
    setPillorders(res);
    console.log(res);
  };
  getPillorders();
}, [loading]);
  return (
    <Page theme={pageTheme.tool}>

<Header title="แสดงข้อมูล" >
<Link component={RouterLink} to="/home">
              <Button
                variant="contained"
                color="default"
              >
                กลับ
              </Button>
              </Link>
    </Header>
<Content>
    <TableContainer component={Paper}>
    <Table className={classes.table} aria-label="simple table">
           <TableHead>
             <TableRow>
               <TableCell align="center">ลำดับ</TableCell>
               <TableCell align="center">หมายเลขใบสั่งยา</TableCell>
               <TableCell align="center">ผู้ป่วย</TableCell>
               <TableCell align="center">ยา</TableCell>
               <TableCell align="center">ทันตแพทย์</TableCell>
               <TableCell align="center">พนักงาน</TableCell>
               <TableCell align="center">วันที่สั่งยา</TableCell>
             </TableRow>
           </TableHead>
           <TableBody>
             {pillorders.map((item: any) =>  (
               <TableRow key={item.id}>
                 <TableCell align="center">{item.id}</TableCell>
                 <TableCell align="center">{item.pillorderNameID}</TableCell>
                 <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
                 <TableCell align="center">{item.edges?.pillorderitem?.pillorderItemName}</TableCell>
                 <TableCell align="center">{item.edges?.dentist?.dentistName}</TableCell>
                 <TableCell align="center">{item.edges?.employee?.employeeName}</TableCell>
                 <TableCell align="center">{moment(item.pillorderDate).format('DD/MM/YYYY HH:mm')}</TableCell>
               </TableRow>
             ))}
           </TableBody>
         </Table>
       </TableContainer>
       </Content>
       </Page>
  );
};

export default ShowPillorder;
