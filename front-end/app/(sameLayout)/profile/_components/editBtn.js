import React from 'react'
import styles from "./profileHeader.module.css"

const EditBtn = ({setProfileNav}) => {
  return (
    <button className={styles.editProfileBtn} onClick={()=>{
      setProfileNav("about")
    }}> Privacy</button>
  )
}

export default EditBtn
