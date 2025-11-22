import React, { useState } from 'react';
import { useQuery, useMutation, gql } from '@apollo/client';
import './App.css';

const GET_DATA = gql`
  query GetData {
    users {
      id
      name
      email
    }
    posts {
      id
      title
      content
      user {
        name
      }
    }
  }
`;

const CREATE_USER = gql`
  mutation CreateUser($name: String!, $email: String!) {
    createUser(input: { name: $name, email: $email }) {
      id
      name
      email
    }
  }
`;

const CREATE_POST = gql`
  mutation CreatePost($title: String!, $content: String!, $userId: String!) {
    createPost(input: { title: $title, content: $content, userId: $userId }) {
      id
      title
      content
      user {
        name
      }
    }
  }
`;

function App() {
  const { loading, error, data, refetch } = useQuery(GET_DATA);
  const [createUser] = useMutation(CREATE_USER, { onCompleted: () => refetch() });
  const [createPost] = useMutation(CREATE_POST, { onCompleted: () => refetch() });

  const [userName, setUserName] = useState('');
  const [userEmail, setUserEmail] = useState('');
  const [postTitle, setPostTitle] = useState('');
  const [postContent, setPostContent] = useState('');
  const [postUserId, setPostUserId] = useState('');

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const handleCreateUser = (e: React.FormEvent) => {
    e.preventDefault();
    createUser({ variables: { name: userName, email: userEmail } });
    setUserName('');
    setUserEmail('');
  };

  const handleCreatePost = (e: React.FormEvent) => {
    e.preventDefault();
    createPost({ variables: { title: postTitle, content: postContent, userId: postUserId } });
    setPostTitle('');
    setPostContent('');
    setPostUserId('');
  };

  return (
    <div className="App">
      <h1>Distributed System Demo</h1>

      <div className="container">
        <div className="section">
          <h2>Users</h2>
          <ul>
            {data.users.map((user: any) => (
              <li key={user.id}>
                {user.name} ({user.email}) - ID: {user.id}
              </li>
            ))}
          </ul>

          <h3>Create User</h3>
          <form onSubmit={handleCreateUser}>
            <input
              type="text"
              placeholder="Name"
              value={userName}
              onChange={(e) => setUserName(e.target.value)}
              required
            />
            <input
              type="email"
              placeholder="Email"
              value={userEmail}
              onChange={(e) => setUserEmail(e.target.value)}
              required
            />
            <button type="submit">Create User</button>
          </form>
        </div>

        <div className="section">
          <h2>Posts</h2>
          <ul>
            {data.posts.map((post: any) => (
              <li key={post.id}>
                <strong>{post.title}</strong> by {post.user.name}
                <p>{post.content}</p>
              </li>
            ))}
          </ul>

          <h3>Create Post</h3>
          <form onSubmit={handleCreatePost}>
            <input
              type="text"
              placeholder="Title"
              value={postTitle}
              onChange={(e) => setPostTitle(e.target.value)}
              required
            />
            <input
              type="text"
              placeholder="Content"
              value={postContent}
              onChange={(e) => setPostContent(e.target.value)}
              required
            />
            <select
              value={postUserId}
              onChange={(e) => setPostUserId(e.target.value)}
              required
            >
              <option value="">Select User</option>
              {data.users.map((user: any) => (
                <option key={user.id} value={user.id}>
                  {user.name}
                </option>
              ))}
            </select>
            <button type="submit">Create Post</button>
          </form>
        </div>
      </div>
    </div>
  );
}

export default App;
