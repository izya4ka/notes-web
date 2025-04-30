import React, { FC, ReactNode, useContext } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';

interface Props {
  children: ReactNode;
}

export const ProtectedRoute: FC<Props> = ({ children }) => {
  const auth = useContext(AuthContext);
  if (!auth) throw new Error('AuthContext not provided');

  return auth.token ? <>{children}</> : <Navigate to="/login" replace />;
};
