// URLDisplay.js
import React, { useEffect, useState } from "react";
import Axios from "axios";
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import FileCopyIcon from '@mui/icons-material/FileCopy';

function URLDisplay({ urlId, onNewUrlId }) {
    const [retrievedUrl, setRetrievedUrl] = useState("");
    const [isCopied, setIsCopied] = useState(false);

    useEffect(() => {
      // Fetch the URL data from the database when the component mounts
      const fetchData = async () => {
        try {
          const response = await Axios.get(`http://localhost:7777/url/${urlId}`);
          const data = response.data;
  
          if (data.url) {
            setRetrievedUrl(data.url);
          }
        } catch (error) {
          console.error("Error while fetching URL data:", error);
        }
      };
  
      fetchData();
    }, [urlId]);

    const handleCopyClick = () => {
        navigator.clipboard.writeText(completeUrl)
          .then(() => {
            setIsCopied(true);
            setTimeout(() => setIsCopied(false), 2000); // Reset isCopied after 2 seconds
          })
          .catch((error) => console.error("Error copying URL:", error));
      };
      const completeUrl = `http://localhost:7777/${retrievedUrl}`;


  return (
    <div>
      <TextField
        label="Shortened URL"
        variant="outlined"
        value={completeUrl}
        fullWidth
        InputProps={{
          endAdornment: (
            <Button
              onClick={handleCopyClick}
              variant="contained"
              color="primary"
              size="medium"
              startIcon={<FileCopyIcon />}
            >
              {isCopied ? 'Copied!' : 'Copy'}
            </Button>
          ),
        }}
      />
    </div>
  );
}

export default URLDisplay;
