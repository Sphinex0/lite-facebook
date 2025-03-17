'use client';

import Link from 'next/link';
import GroupInfo from './_components/groupInfo';
import styles from './groupe.module.css';
import { useEffect, useState } from 'react';

const Groups = () => {
  const [compContent, setCompContent] = useState(null); // Store data directly
  const [loading, setLoading] = useState(false); // For loading state
  const [section , setSection] = useState("all groups")

  // Function to handle the link click (fetch data)
  const handleLinkClick = async (content) => {
    setLoading(true); // Show loading spinner or text
    try {
      const response = await fetch(`http://localhost:8080/api/` + content, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const data = await response.json();
      console.log(data);

      setCompContent(data); 
    } catch (error) {
      console.error('Fetch error:', error);
      setCompContent(null); 
    } finally {
      setLoading(false);
    }
  };

  const renderTable = (data) => {
    if (!data || data.length === 0) {
      return <p>No data available</p>;
    }

    return (
      <div className='feeds'>
        {data.map((row) => (
                    <Link href={`/groups/${row.id}`} key={row.id}><GroupInfo  groupInfo={row}/></Link>
                  ))}
      </div>
      // <table className={styles.table}>
      //   <thead>
      //     <tr>
      //       <th>Title</th>
      //       <th>Description</th>
          
          
        
      //     </tr>
      //   </thead>
      //   <tbody>
      //     {data.map((row, index) => (
      //       <tr key={index}>

      //         <td>{row.title}</td>
      //         <td>{row.description}</td>

             
      //       </tr>
      //     ))}
      //   </tbody>
      // </table>
    );
  };

  useEffect(()=>{
    handleLinkClick("groups")
  },[])

  return (
    <div className='groups'>
      <div className="category">
        <h6 className={section === "all groups" ? "active":""} 
        onClick={() => {
          setSection("all groups")
          handleLinkClick('groups')
        }}>All Groups</h6>
        <h6 className={section === "your groups" ? "active":""} 
        onClick={() => {
          setSection("your groups")
          handleLinkClick('members')
        }}>Your Groups</h6>
      </div>
      {/* <div className={styles.typegroupes}>
        <a href="#!" onClick={() => handleLinkClick('groups')}>
          <div>All Groups</div>
        </a>
        <a href="#!" onClick={() => handleLinkClick('members')}>
          <div>Your Groups</div>
        </a>
      </div> */}

      <div className={styles.comp}>
        {loading && <p>Loading...</p>}
        {compContent && renderTable(compContent)}
      </div>
    </div>
  );
};

export default Groups;
