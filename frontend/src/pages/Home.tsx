import React, { FC, useState, useCallback, KeyboardEvent, useRef, useEffect, MouseEvent } from 'react'
import Grid from '@mui/material/Grid'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Container from '@mui/material/Container'
import AddIcon from '@mui/icons-material/Add'
import ArrowUpwardIcon from '@mui/icons-material/ArrowUpward'
import Tooltip from '@mui/material/Tooltip'
import Avatar from '@mui/material/Avatar'

import Page from '../components/system/Page'
import Row from '../components/widgets/Row'
import SessionTypeButton from '../components/create/SessionTypeButton'
import ModelPicker from '../components/create/ModelPicker'
import ExamplePrompts from '../components/create/ExamplePrompts'
import LoadingSpinner from '../components/widgets/LoadingSpinner'
import { ISessionType, SESSION_TYPE_TEXT } from '../types'

import useLightTheme from '../hooks/useLightTheme'
import useIsBigScreen from '../hooks/useIsBigScreen'
import useRouter from '../hooks/useRouter'
import useSnackbar from '../hooks/useSnackbar'
import useSessions from '../hooks/useSessions'
import useApps from '../hooks/useApps'
import useAccount from '../hooks/useAccount'
import { useStreaming } from '../contexts/streaming'

import {
  SESSION_MODE_FINETUNE,
} from '../types'

const getTimeAgo = (date: Date) => {
  const now = new Date()
  const seconds = Math.floor((now.getTime() - date.getTime()) / 1000)
  const minutes = Math.floor(seconds / 60)
  const hours = Math.floor(minutes / 60)
  const days = Math.floor(hours / 24)

  if (days > 0) return `${days} days ago`
  if (hours > 0) return `${hours} hours ago`
  if (minutes > 0) return `${minutes} minutes ago`
  return 'just now'
}

const Home: FC = () => {
  const isBigScreen = useIsBigScreen()
  const lightTheme = useLightTheme()
  const router = useRouter()
  const snackbar = useSnackbar()
  const sessions = useSessions()
  const account = useAccount()
  const apps = useApps()
  const { NewInference } = useStreaming()
  const [currentPrompt, setCurrentPrompt] = useState('')
  const [currentType, setCurrentType] = useState<ISessionType>(SESSION_TYPE_TEXT)
  const [currentModel, setCurrentModel] = useState<string>('')
  const [loading, setLoading] = useState(false)
  const textareaRef = useRef<HTMLTextAreaElement>(null)

  // Focus textarea on mount
  useEffect(() => {
    if (textareaRef.current) {
      textareaRef.current.focus()
    }
  }, [])

  // Focus textarea when prompt changes (e.g. from example prompts)
  useEffect(() => {
    if (textareaRef.current) {
      textareaRef.current.focus()
    }
  }, [currentPrompt])

  useEffect(() => {
    if(!account.user) return
    apps.loadData()
  }, [
    account.user,
  ])

  const submitPrompt = async () => {
    if (!currentPrompt.trim()) return
    setLoading(true)
    try {
      const session = await NewInference({
        type: currentType,
        message: currentPrompt,
        modelName: currentModel,
      })
      if (!session) return
      await sessions.loadSessions()
      setLoading(false)
      router.navigate('session', { session_id: session.id })
    } catch (error) {
      console.error('Error in submitPrompt:', error)
      snackbar.error('Failed to start inference')
      setLoading(false)
    }
  }

  const openApp = async (appId: string) => {
    router.navigate('new', { app_id: appId });
  }

  const onCreateNewApp = async () => {
    if (!account.user) {
      account.setShowLoginWindow(true)
      return
    }
    const newApp = await apps.createEmptyHelixApp()
    if(!newApp) return false
    apps.loadData()
    router.navigate('app', {
      app_id: newApp.id,
    })
  }

  const handleKeyDown = (e: KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      submitPrompt()
    }
  }

  return (
    <Page
      showTopbar={ isBigScreen ? false : true }
    >
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          minHeight: '100%',
        }}
      >
        {/* Main content */}
        <Box
          sx={{
            flex: 1,
          }}
        >
          <Container
            maxWidth="md"
            sx={{
              py: 4,
              display: 'flex',
              px: { xs: 1, sm: 2, md: 3 },
              overflow: 'hidden',
            }}
          >
            <Grid container spacing={1} justifyContent="center">
              <Grid item xs={12} sx={{ textAlign: 'center', maxWidth: '100%', overflow: 'hidden' }}>
                <Row
                  sx={{
                    display: 'flex',
                    flexDirection: 'row',
                    alignItems: 'center',
                    justifyContent: 'center',
                  }}
                >
                  <Typography
                    sx={{
                      color: '#fff',
                      fontSize: '1.5rem',
                      fontWeight: 'bold',
                      textAlign: 'center',
                      mb: 2,
                    }}
                  >
                    How can I help?
                  </Typography>
                </Row>
                <Row>
                  <Box
                    sx={{
                      width: '100%',
                      border: '1px solid rgba(255, 255, 255, 0.2)',
                      borderRadius: '12px',
                      backgroundColor: 'rgba(255, 255, 255, 0.05)',
                      p: 2,
                      mb: 2,
                    }}
                  >
                    {/* Top row - Chat with Helix */}
                    <Box
                      sx={{
                        display: 'flex',
                        alignItems: 'center',
                        mb: 2,
                      }}
                    >
                      <textarea
                        ref={textareaRef}
                        value={currentPrompt}
                        onChange={(e) => setCurrentPrompt(e.target.value)}
                        onKeyDown={handleKeyDown}
                        rows={2}
                        style={{
                          width: '100%',
                          backgroundColor: 'transparent',
                          border: 'none',
                          color: '#fff',
                          opacity: 0.7,
                          resize: 'none',
                          outline: 'none',
                          fontFamily: 'inherit',
                          fontSize: 'inherit',
                        }}
                        placeholder="Chat with Helix"
                      />
                    </Box>

                    {/* Bottom row - Split into left and right sections */}
                    <Box
                      sx={{
                        display: 'flex',
                        justifyContent: 'space-between',
                        alignItems: 'center',
                        flexWrap: { xs: 'wrap', sm: 'nowrap' },
                        gap: 1,
                      }}
                    >
                      {/* Left section - Will contain SessionTypeButton, ModelPicker and plus button */}
                      <Box
                        sx={{
                          display: 'flex',
                          alignItems: 'center',
                          gap: 1,
                          flexWrap: { xs: 'wrap', sm: 'nowrap' },
                          flex: 1,
                          minWidth: 0,
                        }}
                      >
                        <SessionTypeButton 
                          type={currentType}
                          onSetType={setCurrentType}
                        />
                        <ModelPicker
                          type={currentType}
                          model={currentModel}
                          provider={undefined}
                          displayMode="short"
                          border
                          compact
                          onSetModel={setCurrentModel}
                        />
                        {/* Plus button */}
                        <Tooltip title="Add Documents" placement="top">
                          <Box 
                            sx={{ 
                              width: 32, 
                              height: 32,
                              display: 'flex',
                              alignItems: 'center',
                              justifyContent: 'center',
                              cursor: 'pointer',
                              border: '2px solid rgba(255, 255, 255, 0.7)',
                              borderRadius: '50%',
                              '&:hover': {
                                borderColor: 'rgba(255, 255, 255, 0.9)',
                                '& svg': {
                                  color: 'rgba(255, 255, 255, 0.9)'
                                }
                              }
                            }}
                            onClick={() => {
                              router.navigate('new', {
                                model: currentModel,
                                type: currentType,
                                mode: SESSION_MODE_FINETUNE,
                                rag: true,
                              })
                            }}
                          >
                            <AddIcon sx={{ color: 'rgba(255, 255, 255, 0.7)', fontSize: '20px' }} />
                          </Box>
                        </Tooltip>
                      </Box>

                      {/* Right section - Up arrow icon */}
                      <Box>
                        <Tooltip title="Send Prompt" placement="top">
                          <Box 
                            onClick={submitPrompt}
                            sx={{ 
                              width: 32, 
                              height: 32,
                              display: 'flex',
                              alignItems: 'center',
                              justifyContent: 'center',
                              cursor: loading ? 'default' : 'pointer',
                              border: '1px solid rgba(255, 255, 255, 0.7)',
                              borderRadius: '8px',
                              opacity: loading ? 0.5 : 1,
                              '&:hover': loading ? {} : {
                                borderColor: 'rgba(255, 255, 255, 0.9)',
                                '& svg': {
                                  color: 'rgba(255, 255, 255, 0.9)'
                                }
                              }
                            }}
                          >
                            {loading ? (
                              <LoadingSpinner />
                            ) : (
                              <ArrowUpwardIcon sx={{ color: 'rgba(255, 255, 255, 0.7)', fontSize: '20px' }} />
                            )}
                          </Box>
                        </Tooltip>
                      </Box>
                    </Box>
                  </Box>
                </Row>
                <Row>
                  <Box
                    sx={{
                      width: '100%',
                      // px: 2,
                      mb: 6,
                    }}
                  >
                    <ExamplePrompts
                      header={false}
                      layout="vertical"
                      type={currentType}
                      onChange={setCurrentPrompt}
                    />
                  </Box>
                </Row>
                <Row
                  sx={{
                    display: 'flex',
                    flexDirection: 'row',
                    alignItems: 'left',
                    justifyContent: 'left',
                    mb: 1,
                  }}
                >
                  Recent Apps
                </Row>
                <Row
                  sx={{
                    display: 'flex',
                    flexDirection: 'row',
                    alignItems: 'left',
                    justifyContent: 'left',
                    mb: 1,
                  }}
                >
                  <Grid container spacing={1} justifyContent="left">
                    {
                      [...apps.data]
                        .sort((a, b) => new Date(b.updated).getTime() - new Date(a.updated).getTime())
                        .map((app) => (
                          <Grid item xs={12} sm={6} md={4} lg={4} xl={4} sx={{ textAlign: 'left', maxWidth: '100%' }} key={ app.id }>
                            <Box
                              sx={{
                                borderRadius: '12px',
                                border: '1px solid rgba(255, 255, 255, 0.2)',
                                p: 1.5,
                                pb: 0.5,
                                cursor: 'pointer',
                                '&:hover': {
                                  backgroundColor: 'rgba(255, 255, 255, 0.05)',
                                },
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'flex-start',
                                gap: 1,
                                width: '100%',
                                minWidth: 0,
                              }}
                              onClick={() => openApp(app.id)}
                            >
                              <Avatar
                                sx={{
                                  width: 28,
                                  height: 28,
                                  backgroundColor: 'rgba(255, 255, 255, 0.1)',
                                  color: '#fff',
                                  fontWeight: 'bold',
                                  border: (theme) => app.config.helix.avatar ? '2px solid rgba(255, 255, 255, 0.8)' : 'none',
                                }}
                                src={app.config.helix.avatar}
                              >
                                {app.config.helix.name[0].toUpperCase()}
                              </Avatar>
                              <Box sx={{ textAlign: 'left', width: '100%', minWidth: 0 }}>
                                <Typography sx={{ 
                                  color: '#fff',
                                  fontSize: '0.95rem',
                                  lineHeight: 1.2,
                                  fontWeight: 'bold',
                                  overflow: 'hidden',
                                  textOverflow: 'ellipsis',
                                  whiteSpace: 'nowrap',
                                  width: '100%',
                                }}>
                                  { app.config.helix.name }
                                </Typography>
                                <Typography variant="caption" sx={{ 
                                  color: 'rgba(255, 255, 255, 0.5)',
                                  fontSize: '0.8rem',
                                  lineHeight: 1.2,
                                }}>
                                  { getTimeAgo(new Date(app.updated)) }
                                </Typography>
                              </Box>
                            </Box>
                          </Grid>
                        ))
                    }
                    <Grid item xs={12} sm={6} md={4} lg={4} xl={4} sx={{ textAlign: 'center' }}>
                      <Box
                        sx={{
                          borderRadius: '12px',
                          border: '1px dashed rgba(255, 255, 255, 0.2)',
                          p: 1.5,
                          pb: 0.5,
                          cursor: 'pointer',
                          '&:hover': {
                            backgroundColor: 'rgba(255, 255, 255, 0.05)',
                          },
                          display: 'flex',
                          flexDirection: 'column',
                          alignItems: 'flex-start',
                          gap: 1,
                        }}
                        onClick={() => onCreateNewApp()}
                      >
                        <Box
                          sx={{
                            width: 28,
                            height: 28,
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center',
                            borderRadius: '50%',
                            backgroundColor: 'rgb(0, 153, 255)',
                          }}
                        >
                          <AddIcon sx={{ color: '#fff', fontSize: '20px' }} />
                        </Box>
                        <Box sx={{ textAlign: 'left' }}>
                          <Typography sx={{ 
                            color: '#fff',
                            fontSize: '0.95rem',
                            lineHeight: 1.2,
                            fontWeight: 'bold',
                          }}>
                            Create new app
                          </Typography>
                          <Typography variant="caption" sx={{ 
                            color: 'rgba(255, 255, 255, 0.5)',
                            fontSize: '0.8rem',
                            lineHeight: 1.2,
                          }}>
                            &nbsp;
                          </Typography>
                        </Box>
                      </Box>
                    </Grid>
                  </Grid>
                </Row>
              </Grid>
            </Grid>
          </Container>
        </Box>

        {/* Footer */}
        <Box
          component="footer"
          sx={{
            py: 2,
            mt: 'auto',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            borderTop: (theme) => `1px solid ${theme.palette.divider}`,
          }}
        >
          <Typography
            sx={{
              color: lightTheme.textColorFaded,
              fontSize: '0.8rem',
            }}
          >
            Open source models can make mistakes. Check facts, dates and events.
          </Typography>
        </Box>
      </Box>
    </Page>
  )
}

export default Home