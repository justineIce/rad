`import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/manage',
  name: 'manage',
  meta,
  redirect: { name: '{{range $index, $item := .tableList}}{{if eq $index 0}}{{$item.SingName}}{{end}}{{end}}' },
  component: layoutHeaderAside,
  children:  [
    {{ range .tableList }}{ path: '{{.SingName}}', name: '{{.SingName}}', component: () => import('@/pages/{{.TableName}}'), meta: { ...meta, title: '{{.TableRemark}}' } },
    {{end}}
  ]
}`
