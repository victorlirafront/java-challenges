import React, { createContext, useContext, useState, ReactNode } from 'react';

interface AsideMenuContextType {
  displayAsideMenu: boolean;
  toggleAsideMenu: () => void;
  closeMenu: () => void;
}

interface AsideMenuProviderProps {
  children: ReactNode;
}

const AsideMenuContext = createContext<AsideMenuContextType | undefined>(undefined);

export const AsideMenuProvider: React.FC<AsideMenuProviderProps> = ({ children }) => {
  const [displayAsideMenu, setDisplayAsideMenu] = useState<boolean>(false);

  const toggleAsideMenu = () => {
    setDisplayAsideMenu(prev => !prev);
  };

  const closeMenu = () => {
    setDisplayAsideMenu(false)
  }

  return (
    <AsideMenuContext.Provider value={{ displayAsideMenu, toggleAsideMenu, closeMenu }}>
      {children}
    </AsideMenuContext.Provider>
  );
};

export const useAsideMenu = (): AsideMenuContextType => {
  const context = useContext(AsideMenuContext);
  if (!context) {
    throw new Error('useAsideMenu must be used within an AsideMenuProvider');
  }
  return context;
};
