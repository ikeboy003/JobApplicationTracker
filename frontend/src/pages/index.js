import { useState, useEffect } from "react";
import axios from "axios";
import Link from 'next/link';

export default function Jobs() {
  const [jobs, setJobs] = useState([]);
  const [newJobName, setNewJobName] = useState("");
  const [newJobCompany, setNewJobCompany] = useState("");
  const [newAppStatus, setNewAppStatus] = useState("applied"); 

  useEffect(() => {
    async function fetchJobs() {
      try {
        
        const response = await axios.get("http://localhost:3000/listAll");
        setJobs(response.data);
      } catch (error) {
        console.error("Error fetching jobs:", error);
      }
    }
  
    fetchJobs();
  }, []);

  const handleSubmit = async (e) => {
    

    try {
      const response = await axios.post("http://localhost:3000/newJob", {
        jobName: newJobName,
        jobCompany: newJobCompany,
        appStatus: newAppStatus,
      });

     
      setJobs([...jobs, response.data]);

      
      setNewJobName("");
      setNewJobCompany("");
      setNewAppStatus("applied");
    } catch (error) {
      console.error("Error adding job:", error);
    }
  };

  return (
    <div>
      <h1>Job Applied</h1>
      <table>
        <thead>
          <tr>
            <th>Job ID</th>
            <th>Job Name</th>
            <th>Job Company</th>
            <th>Application Status</th>
          </tr>
        </thead>
        <tbody>
          {jobs.map((job) => (
            <tr key={job.jobID}>
              <td>{job.jobID}</td>
              <td>{job.jobName}</td>
              <td>{job.jobCompany}</td>
              <td>{job.appStatus}</td>
            </tr>
          ))}
        </tbody>
      </table>
        <form onSubmit={handleSubmit}>
        <label>
          Job Name:
          <input
            type="text"
            value={newJobName}
            onChange={(e) => setNewJobName(e.target.value)}
          />
        </label>
        <br />
        <label>
          Job Company:
          <input
            type="text"
            value={newJobCompany}
            onChange={(e) => setNewJobCompany(e.target.value)}
          />
        </label>
        <br />
        <label>
          Application Status:
          <select
            value={newAppStatus}
            onChange={(e) => setNewAppStatus(e.target.value)}
          >
            <option value="applied">Applied</option>
            <option value="rejected">Rejected</option>
            <option value="interviewed">Interviewed</option>
          </select>
        </label>
        <br />
        <button type="submit">Add Job</button>
      </form>

      <Link href="/addFile">
      <button>Add a File</button>
      </Link>
    </div>
  );
}
