import React from 'react';
import './Home.css';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    root: {
        '& > *': {
            margin: theme.spacing(1),
        },
    },
    table: {
        minWidth: 650,
    },
    button: {
        margin: theme.spacing(1),
    },
}));

function Home() {
    const classes = useStyles();
    return (
        <div className={classes.root}>
            <h1>Hello World</h1>
        </div>
    );
}

export default Home;