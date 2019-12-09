Name:           few_latest_artefact_github_release
Version:        0.1
Release:        1
Summary:        Get a few latest artefact from github release.
License:        ASL 2.0 
Source0:        main.go
Source1:        go.mod
Source2:        go.sum
BuildRequires:  golang

%description
Get a few latest artefact from github release.

%build
export GOPATH=%{_builddir}/_build
# go get github.com/google/go-github/v28/github
mkdir -p _build/src/github.com/patsevanton/
git clone https://github.com/patsevanton/few_latest_artefact_github_release.git $GOPATH/src/github.com/patsevanton/few_latest_artefact_github_release
git config http.postBuffer 524288000
ls
pwd
pushd _build/src/github.com/patsevanton/few_latest_artefact_github_release
ls
pwd
go build -o ../../../../../few_latest_artefact_github_release
ls
pwd
popd

%install
install -d %{buildroot}%{_bindir}
install -p -m 0755 ./few_latest_artefact_github_release %{buildroot}%{_bindir}/few_latest_artefact_github_release

%files
%defattr(-,root,root,-)
%{_bindir}/few_latest_artefact_github_release
