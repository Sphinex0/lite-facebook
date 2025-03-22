'use client';

import Link from 'next/link';
import GroupInfo from './_components/groupInfo';
import './groupe.css';
import { useEffect, useState } from 'react';
import { Add, DisabledByDefault, TurnSharpLeft } from '@mui/icons-material';
import { useRouter } from 'next/navigation';
import { FetchApi } from '@/app/helpers';
let type = "groups"

const Groups = () => {
  const [compContent, setCompContent] = useState(null);
  const [loading, setLoading] = useState(false);
  const [section, setSection] = useState("all groups")
  const [groupCreated, setGroupCreated] = useState(false);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [image, setImage] = useState('');
  const redirect = useRouter()

  const handleSubmit = async (e) => {
    e.preventDefault();

    const formDataToSend = new FormData();
    formDataToSend.append('Title', title);
    formDataToSend.append('Description', description);
    formDataToSend.append('image', image);

    console.log(formDataToSend);

    setTitle("")
    setDescription("")
    setImage(null)


    try {
      const response = await FetchApi(`/api/groups/store`, redirect , {
        method: 'POST',
        body: formDataToSend
      });


      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      handleLinkClick(type)



    } catch (error) {
      console.error('Fetch error:', error);
      setCompContent(null);
    } finally {
      setLoading(false);
    }



  }



  const CreateGroup = () => {
    console.log("groupCreated", groupCreated);

    setGroupCreated(true)
    const element = document.querySelector('.formclass')
    element.style.display = "flex"
  }

  const SeeClick = () => {
    if (groupCreated == true) {
      setGroupCreated(false)
      const element = document.querySelector('.formclass')
      element.style.display = "none"
    }

  }

  const handleLinkClick = async (content) => {
    type = content
    setLoading(true); // Show loading spinner or text
    try {
      const response = await FetchApi(`/api/` + content,redirect ,{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
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
          <Link href={`/groups/${row.id}`} key={row.id}><GroupInfo groupInfo={row} /></Link>
        ))}
      </div>
    );
  };

  useEffect(() => {
    handleLinkClick(type)
  }, [])

  return (
    <div className='groups'>
      <div className='layout' ></div>
      <div className="category">
        <h6 className={section === "all groups" ? "active" : ""}
          onClick={() => {
            setSection("all groups")
            handleLinkClick('groups')
          }}>All Groups</h6>
        <h6 className={section === "your groups" ? "active" : ""}
          onClick={() => {
            setSection("your groups")
            handleLinkClick("members")
          }}>Your Groups</h6>
      </div>

      <div>
        <button className='create'
          onClick={() => {
            CreateGroup()
          }}>
          <Add />
        </button>
      </div>

      <div className='formclass'>
        <div onClick={() => { SeeClick() }}>
          <DisabledByDefault />
        </div>
        <form method='POST' aria-multiselectable onSubmit={handleSubmit}>
          <label htmlFor='title' >title</label>
          <input type='text' className='title' id='title'   value={title} onChange={(e) => setTitle(e.target.value)} />
          <label htmlFor='descriptopn'>descriptopn</label>
          <input type='text' className='descriptopn' id='descriptopn' value={description}  onChange={(e) => setDescription(e.target.value)} />
          <label htmlFor='image'>image</label>
          <input type='file' className='image' id='image' value={image} onChange={(e) => setImage(e.target.value)} />
          <button className='button' type='onSubmit'> Submit </button>
        </form>
      </div>

      <div className='comp'>
        {loading && <p>Loading...</p>}
        {compContent && renderTable(compContent)}
      </div>
    </div>
  );
};

export default Groups;
