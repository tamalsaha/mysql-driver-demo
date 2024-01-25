package main

import "testing"

func TestCanonicalMySQLDSN(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		//{
		//	name:    "missing port",
		//	args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value"},
		//	want:    "mysql-demo.mysql.svc",
		//	wantErr: false,
		//},
		//{
		//	name:    "default port",
		//	args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value"},
		//	want:    "mysql-demo.mysql.svc",
		//	wantErr: false,
		//},
		//{
		//	name:    "custom port",
		//	args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:1234)/dbname?param=value"},
		//	want:    "mysql-demo.mysql.svc",
		//	wantErr: false,
		//},
		// "tcp://mysql-demo.mysql.svc:3306"
		{
			name:    "custom port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:3306"},
			want:    "tcp(mysql-demo.mysql.svc:3306)/",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:3306/dbname"},
			want:    "tcp(mysql-demo.mysql.svc:3306)/dbname",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:3306/dbname?param=value"},
			want:    "tcp(mysql-demo.mysql.svc:3306)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "tcp://username:password@mysql-demo.mysql.svc:3306/dbname?param=value"},
			want:    "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CanonicalMySQLDSN(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("CanonicalMySQLDSN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CanonicalMySQLDSN() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMySQLHost(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "missing port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:1234)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMySQLHost(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMySQLHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMySQLHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMySQLPort(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "missing port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value"},
			want:    "3306",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value"},
			want:    "3306",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:1234)/dbname?param=value"},
			want:    "1234",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMySQLPort(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMySQLPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMySQLPort() got = %v, want %v", got, tt.want)
			}
		})
	}
}
