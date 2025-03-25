import { Lock } from '@mui/icons-material'
import React from 'react'
import styles from"./notAllowed.module.css"
const NotAllowed = () => {
  return (
    <div className={styles.container}>
      <Lock/>
      <span>join / follow to see</span>
    </div>
  )
}

export default NotAllowed
