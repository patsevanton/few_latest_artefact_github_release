Name:           change-name-rpm-kubernetes
Version:        0.1
Release:        3%{?dist}
Summary:        A simple Golang command-line utility to change name rpm kubernetes repo.
License:        ASL 2.0 
Source0:        main.go
BuildRequires:  golang

%description
A simple Golang command-line utility to change name rpm kubernetes repo.

%build
mkdir -p _build/src/github.com/patsevanton/change-name-rpm-kubernetes-repo
cp ../SOURCES/main.go _build/src/github.com/patsevanton/change-name-rpm-kubernetes-repo
export GOPATH=$(pwd)/_build
export PATH=$PATH:$(pwd)/_build/bin

pushd _build/src/github.com/patsevanton/change-name-rpm-kubernetes-repo
go build -o ../../../../../change-name-rpm-kubernetes
popd

%install
install -d %{buildroot}%{_bindir}
install -p -m 0755 ./change-name-rpm-kubernetes %{buildroot}%{_bindir}/change-name-rpm-kubernetes

%files
%defattr(-,root,root,-)
%{_bindir}/change-name-rpm-kubernetes
