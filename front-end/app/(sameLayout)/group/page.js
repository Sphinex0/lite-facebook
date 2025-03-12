'use client';

import styles from './groupe.module.css';
import { useState } from 'react';

const Group = () => {
  const [compContent, setCompContent] = useState('');

  const handleLinkClick = (content) => {
    setCompContent(content);
  };

  return (
    <div>
      <div className={styles.typegroupes}>
        <a href="#!" onClick={() => handleLinkClick('All groups content')}>
          <div>all group</div>
        </a>
        <a href="#!" onClick={() => handleLinkClick('All group your members content')}>
          <div>all group your members</div>
        </a>
      </div>
      <div className={styles.comp}> {/* Use CSS module for styling */}
        {compContent}
      </div>
    </div>
  );
};

export default Group;
