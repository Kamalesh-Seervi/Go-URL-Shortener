import * as React from "react";
import { useState } from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import Link from "@mui/material/Link";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import Axios from "axios";
import URLDisplay from "./Redirect";
import Snackbar from '@mui/material/Snackbar'; // Import Snackbar
import MuiAlert from '@mui/material/Alert';
import LinkIcon from '@mui/icons-material/Link';


// import { createTheme } from '@mui/material/styles';



function Copyright(props) {
  return (
    <Typography
      variant="body2"
      color="text.secondary"
      align="center"
      {...props}
    >
      {"Copyright © "}
      <Link color="inherit" href="https://mui.com/">
        Your Website
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

const defaultTheme = createTheme();

// const theme = createTheme({
//   palette: {
//     primary: {
//       main: 'black', // Text color
//     },
//     background: {
//       default: 'white', // Background color
//     },
//   },
// });
export default function SignUp() {
  // const [shortenedUrl, setShortenedUrl] = useState("");
  const [inputUrl, setInputUrl] = useState("");
  const [urlId, setUrlId] = useState("");
  const [error, setError] = useState(null); 

  // const [getUrl,setGetUrl]=useState("") 
  // const [retrievedUrl, setRetrievedUrl] = useState(""); // State to hold the input URL



  const handleSubmit = async (event) => {
    event.preventDefault();

    if (inputUrl === "") {
      setError("URL cannot be empty"); // Set the error message
      return;
    }

    try {
      const response = await Axios.post("http://localhost:7777/url", {
        redirect: inputUrl,
        random: true,
      });

      const newUrlId = response.data.id;
      setUrlId(newUrlId);
    } catch (error) {
      console.error("Error while shortening URL:", error);
    }
  };

  const handleCloseError = () => {
    setError(null); // Clear the error message when Snackbar is closed
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        />
        <Grid item xs={12} noValidate sx={{ mt: 4 }} container justifyContent="center" alignItems="center">
        <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
  <LinkIcon />
</Avatar>
        </Grid>
        <Grid item xs={12} container justifyContent="center" alignItems="center">
        <Typography component="h1" variant="h5">
          ShortDash URL Shortner
        </Typography>
        </Grid>


    
        <Box component="form" noValidate sx={{ mt: 8 }}>
          
          <Grid container spacing={2}>
            
            <Grid item xs={12} container justifyContent="center" alignItems="center">
              <TextField
                fullWidth
                id="url"
                label="Enter a URL"
                value={inputUrl}
                onChange={(e) => setInputUrl(e.target.value)} // Update inputUrl state on change
                name="url" // Add this line
              />
              <Button
                type="button"
                onClick={handleSubmit}
                variant="contained"
                color="primary"
                style={{ marginTop: "16px" }}
              >
                Generate URL
              </Button>
            </Grid>
            <Grid item xs={12}>
            <URLDisplay urlId={urlId} onNewUrlId={(newUrlId) => setUrlId(newUrlId)} />
</Grid>
          </Grid>
          <Snackbar open={!!error} autoHideDuration={4000} onClose={handleCloseError} anchorOrigin={{ vertical: 'top', horizontal: 'center'}}>
        <MuiAlert onClose={handleCloseError} severity="error">
          {error}
        </MuiAlert>
      </Snackbar>


        </Box>
        <Copyright sx={{ mt: 5 }} />
      </Container>
    </ThemeProvider>
  );
}
