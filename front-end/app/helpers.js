import { useEffect, useRef } from "react";

async function Checkuservalidity() {
    try {
        const response = await fetch("/checkuser")
        return response.json()
    } catch (err) {
        return false
    }
}

const pluralize = (number, unit) => {
    if (number === 1) {
        return `${number} ${unit} ago`;
    } else {
        return `${number} ${unit}s ago`;
    }
}
export const timeAgo = (unixTimestamp) => {
    // Convert Unix timestamp (in seconds) to a Date object (requires milliseconds)
    const pastDate = new Date(unixTimestamp * 1000);
    const currentDate = new Date();

    // Calculate the difference in milliseconds
    const diffMs = currentDate - pastDate;

    // Handle future timestamps
    if (diffMs < 0) {
        return "in the future";
    }

    // Convert difference to seconds
    const diffSeconds = Math.floor(diffMs / 1000);

    // Define time thresholds and calculate appropriate unit
    if (diffSeconds < 60) {
        // Less than 1 minute
        if (diffSeconds < 10) {
            return "just now"; // Very recent, less than 10 seconds
        } else {
            return pluralize(diffSeconds, "second");
        }
    } else if (diffSeconds < 3600) {
        // Less than 1 hour (3600 seconds)
        const minutes = Math.floor(diffSeconds / 60);
        return pluralize(minutes, "minute");
    } else if (diffSeconds < 86400) {
        // Less than 1 day (86400 seconds)
        const hours = Math.floor(diffSeconds / 3600);
        return pluralize(hours, "hour");
    } else if (diffSeconds < 2592000) {
        // Less than 30 days (2592000 seconds)
        const days = Math.floor(diffSeconds / 86400);
        return pluralize(days, "day");
    } else if (diffSeconds < 31536000) {
        // Less than 1 year (31536000 seconds, approximating 365 days)
        const months = Math.floor(diffSeconds / 2592000); // Approximate month as 30 days
        return pluralize(months, "month");
    } else {
        // 1 year or more
        const years = Math.floor(diffSeconds / 31536000); // Approximate year as 365 days
        return pluralize(years, "year");
    }
}


export const likeArticle = async (like, article_id, setLikes, setDislikes, likeState, setLikeState) => {
    try {
        const response = await fetch("http://localhost:8080/api/reactions/store", {
            method: "POST",
            credentials: "include",
            body: JSON.stringify({ like, article_id })
        })

        console.log("status:", response.status)
        if (response.ok) {
            console.log("liked or dislike")
            if (like == 1 && likeState == 1) {
                setLikes((prv) => prv - 1)
                setLikeState(0)
            } else if (like == -1 && likeState == -1) {
                setDislikes((prv) => prv - 1)
                setLikeState(0)
            } else if (like == 1 && likeState == -1) {
                setDislikes((prv) => prv - 1)
                setLikes((prv) => prv + 1)
                setLikeState(1)
            } else if (like == -1 && likeState == 1) {
                setDislikes((prv) => prv + 1)
                setLikes((prv) => prv - 1)
                setLikeState(-1)
            } else if (like == 1 && likeState == 0) {
                setLikes((prv) => prv + 1)
                setLikeState(1)
            } else if (like == -1 && likeState == 0) {
                setDislikes((prv) => prv + 1)
                setLikeState(-1)
            }
        }

    } catch (error) {
        console.log(error)
    }

}

export const addArticle = async(e,setAtricle, {parent, group}) => {
        e.preventDefault()
        try {
          const formData = new FormData(e.target)
          formData.append("group_id",group || 0)
          formData.append("parent",parent || 0)
          console.log(formData)
          const response = await fetch("http://localhost:8080/api/articles/store", {
            method: "POST",
            credentials: "include",
            body: formData
          })
    
          console.log("status:", response.status)
          if (response.ok) {
            const article = await response.json()
            const newArticle = {
              article, // The article data from the server
              user_info: {}, // Empty user_info as before
              likes: 0, // Explicitly set initial likes
              disLikes: 0, // Explicitly set initial dislikes
              comments_count: 0, // Explicitly set initial comments count
              like: 0, // Initial like state (0 = neutral)
            };
            setAtricle((prv) => [newArticle, ...prv]);
            return true
            setModalDisplay(false)
            setContent("")
          }
    
        } catch (error) {
          console.log(error)
        }
      
}

export const useOnVisible = (ref, callback, once = true, threshold = 0.1) => {
    const observerRef = useRef(null);
    
    useEffect(() => {
        console.log("heress",ref)
      if (!ref.current) return;
  
      const observer = new IntersectionObserver(
        ([entry]) => {
          if (entry.isIntersecting) {
            callback();
            if (once) {
              observer.disconnect();
            }
          }
        },
        { threshold }
      );
  
      observer.observe(ref.current);
      observerRef.current = observer; // Store the observer
  
      return () => {
        observer.disconnect(); // Cleanup
      };
    }, [ref, callback, once, threshold]);
  };
