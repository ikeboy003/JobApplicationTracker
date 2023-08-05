import { useState, useEffect } from "react";
import axios from "axios";
import Link from 'next/link';

export default function Jobs() {
  const [jobs, setJobs] = useState([]);

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

      <Link href="/addFile">
      <button>Add a File</button>
      </Link>
    </div>
  );
}
