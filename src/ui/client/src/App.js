import React, { useState} from "react";
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import 'fontsource-roboto';
import GitHubIcon from '@material-ui/icons/GitHub';
import SearchBar from "material-ui-search-bar";
import { Alert } from '@material-ui/lab';
import LinkIcon from '@material-ui/icons/Link';
import 'is-url'
import Typewriter from 'typewriter-effect';


function App() {
  // States for the website
  const [url, setUrl] = useState("");
  const [showResults, setShowResults] = useState(false);
  const [showError, setShowError] = useState(false);
  const [value, setValue] = useState("")

  // On submit of link
  const handleSubmit = (evt) => {
    
    // Check if valid URL
    let check = isValidURL({ value })

    // If true, show generated link else invalid info
    if (check === true) {

      let val = returnStr({ value })
      
      fetch('/', {
        method: 'post',
        body: val,
      })
        .then((response) => response.json())
        .then((responseJson) => {
          setUrl(responseJson["encoded_string"])
          setShowResults(true)
          setShowError(false)
        })
        .catch((error) => {
          console.log(error)
        });
    }

    if (check===false){
      setShowError(true)
    }
    setShowResults(false)

  }

  const useStyles = makeStyles((theme) => ({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    appBar: {
      paddingLeft: 30,
    },
    title: {
      marginLeft: "auto",
      marginRight: 0,
      color: '#000000',
      fontSize: 15,
      fontWeight: 400,
    },
    Maintitle: {
      color: '#000000',
      fontSize: 20,
      fontWeight: 600,
      marginRight: 16,
      marginLeft: -30,
      textDecorationStyle: 'dotted',
    },
    heading: {
      paddingTop: '5%',
      fontWeight: "bold",
      color: '#1D3557',
      [theme.breakpoints.down('md')]: {
        fontSize: 30,
      },
    },
    SearchBar: {
      paddingTop: '5%',
    },
    Alert: {
      paddingTop: '5%',
    },

  }));

  const classes = useStyles();

  // Code snipped for showing link
  const Results = () => (
    <div id="results" className={classes.Alert}>
      <Grid container justify="center">
        <Alert align="center" severity="success">
          {url}
        </Alert>
      </Grid>
      <Grid container justify="center">
        <a className='GitHub' href={url}><LinkIcon></LinkIcon></a>
      </Grid>
    </div>
  )

  // Code snipped for error
  const ErrorCode = () => (
    <div id="results" className={classes.Alert}>
      <Grid container justify="center">
        <Alert align="center" severity="info">
          Invalid URL
          </Alert>
      </Grid>
    </div>
  )



  return (
    <div className={classes.root}>
      <AppBar elevation={0} className={classes.appBar} position="static" style={{ background: '#FFFFFF' }}>
        <Toolbar>
          <Typography variant="h6" className={classes.Maintitle}>
            <span className='titleDe'>Shortify</span>
          </Typography>
          <a className={classes.title} href={"https://github.com/reedkihaddi/URL-Shorten"}>
            <GitHubIcon className="GitHub"></GitHubIcon></a>
        </Toolbar>
      </AppBar>
      <Typography className={classes.heading} variant="h3" gutterBottom align="center">
        Create a <span className="underline">short </span><Typewriter
  options={{
    strings: ['uniform resource locator', 'URL'],
    autoStart: true,
    loop: true,
    wrapperClassName: 'underline-continued'
  }}
/>
      </Typography>
      <div className={classes.SearchBar}>
        <SearchBar
          placeholder="Enter a URL"
          onRequestSearch={handleSubmit}
          onChange={event => {                                 //adding the onChange event
            setValue(event)
          }}
          style={{
            margin: '0 auto',
            maxWidth: '50%',
          }}
        />
        {showResults ? <Results /> : null}
        {showError ? <ErrorCode /> : null}
      </div>
    </div>
  );
}

function returnStr(string) {
  let val = String(string.value)
  val.replace(/\s+/g, '')
  return val
}

function isValidURL(string) {
  let val = String(string.value)
  val.replace(/\s+/g, '')
  try {
    new URL(val);
  } catch (_) {
    return false;
  }

  return true;
};

export default App;
