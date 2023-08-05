import axios from "axios"
import { useState } from 'react';

export default function FileUpload() {
    const [file, setFile] = useState(null);
  
    const handleFileChange = (event) => {
      setFile(event.target.files[0]);
    };
  
    const handleUpload = async () => {
      if (file) {
        const formData = new FormData();
        formData.append('file', file);
  
        try {
          const response = await axios.post('http://localhost:3000/upload', formData);
          console.log('File uploaded:', response.data);
        } catch (error) {
          console.error('Error uploading file:', error);
        }
      }
    };
  
    return (
      <div>
        <input type="file" onChange={handleFileChange} />
        <button onClick={handleUpload}>Upload File</button>
      </div>
    );
  }