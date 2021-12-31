import React, { useState } from 'react';
import './apartments.css';
import Button from '@mui/material/Button';
import { makeStyles } from '@mui/styles';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import { ListApartmentRequest } from '../proto/listings/apartment_pb'
import { getListingsClient } from "../clients";

const srv = getListingsClient()

const useStyles = makeStyles((theme) => ({
    root: {
        '& > *': {
            margin: 1,
        },
    },
    table: {
        minWidth: 650,
    },
    button: {
        margin: 1,
    },
})); 

function ListApartments() {
    // Declare a new state variable, which we'll call "count"
    const [listResult, setListResult] = useState([]);
    const classes = useStyles();

    const LA = () => {
        const listReq = new ListApartmentRequest()
        srv.listApartments(listReq, {}).then((result) => {
            console.log(result.toObject());
            setListResult(result.getApartmentsList())
        })
    }

    return (
        <div className={classes.root}>
            <Typography variant="h6" className={classes.title}>
            <Button onClick={LA} variant="contained" color="primary" type="submit">Get Apartments</Button>
            </Typography>
            <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell align="left">Address</TableCell>
                            <TableCell align="left">Neighborhood</TableCell>
                            <TableCell align="left">Description</TableCell>
                            <TableCell align="left">Rent</TableCell>
                            <TableCell align="left">Sqft</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {listResult.map((apartment) => (
                            < TableRow key = { apartment.getId() }>
                                <TableCell component="th" scope="row">
                                    {apartment.getName()}
                                </TableCell>
                                <TableCell align="left">{apartment.getFullAddress()}</TableCell>
                                <TableCell align="left">{apartment.getNeighborhood()}</TableCell>
                                <TableCell align="left">{apartment.getDescription()}</TableCell>
                                <TableCell align="left">{apartment.getRent()}</TableCell>
                                <TableCell align="left">{apartment.getSqft()}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    );
}

export default ListApartments;