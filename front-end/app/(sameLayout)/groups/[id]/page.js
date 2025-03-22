"use client"
import styles from './group.module.css'
import { useState, useEffect } from 'react';
import { use } from "react";
import Posts from './posts';
import Members from './members';
import joinGroup from "./function";
import Events from "./Events";
import { useRouter } from 'next/navigation';
import { red } from '@mui/material/colors';
import { FetchApi } from '@/app/helpers';
export default function ShowGroup({ params }) {
  const id = use(params).id;

  const [groupNav, setGroupNav] = useState("posts")
  const [groupData, setGroupData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isAllowed, setIsAllowed] = useState(false)
  const [isAction, setIsAction] = useState("")
  

  const redirect = useRouter()
  console.log(JSON.stringify({ id: parseInt(id) }))
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await FetchApi('/api/group', redirect, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ id: parseInt(id) }),
        });
        JSON.stringify({ id })
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }

        const data = await response.json();
        console.log(data);
        setIsAction(data.action)
        setGroupData(data);
        if (data.action === "accepted") {
          setIsAllowed(true)
        } else {
          setIsAllowed(false)
        }
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [id]);

  if (loading) return <h1>Loading...</h1>;
  if (error) return <h1>Error: {error}</h1>;

  return (
    <div>
      <div className={styles.profileHeader}>
        <div className={styles.basicInfo}>
          <div className={styles.p10}>
            <img className={styles.image}></img>
          </div>
          <div className={styles.g2}>

            <h1 className={styles.title}>{groupData.group_info.title}</h1>
            <span className={styles.nickname}>{groupData.group_info.description}</span><br />
            <span className={styles.followText}>{new Date(groupData.group_info.created_at).toLocaleDateString()}</span><br />
          </div>
          <div className={`${styles.g1} ${styles.btnSection}`}>
            {isAllowed
              ? <button className={styles.editProfileBtn}
                onClick={() => {
                  leaveGroup(id)
                }}
              >send Invite</button>


              : <button className={styles.editProfileBtn}
                onClick={() => {
                  joinGroup(id, groupData.group_info.creator, setIsAction, isAction, redirect)
                }}
              >{isAction}</button>
            }

          </div>
        </div>
        <div className={styles.profileNav}>
          <ul className={styles.navUl}>
            <li className={`${styles.navLi} ${groupNav === "posts" ? styles.active : ""}`}
              onClick={() => {
                setGroupNav("posts")
              }}
            >Posts
            </li>

            <li className={`${styles.navLi} ${groupNav === "events" ? styles.active : ""}`}
              onClick={() => {
                setGroupNav("events")
              }}
            >Events</li>

            <li className={`${styles.navLi} ${groupNav === "members" ? styles.active : ""}`}
              onClick={() => {
                setGroupNav("members")
              }}
            >Members</li>
          </ul>

        </div>
        {groupNav == "posts" && <Posts groupID={id} setIsAllowed={setIsAllowed} />}
        {groupNav == "members" && <Members groupID={id} />}
        {groupNav == "events" && <Events groupID={id} />}

        {isAllowed ? "" : "join group first"}


      </div>

    </div>
  );

}
