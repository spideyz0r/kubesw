%global go_version 1.18.10
%global go_release go1.18.10

Name:           kubesw
Version:        0.0.8
Release:        1%{?dist}
Summary:        kubesw, a tool to switch context and namespaces per terminal
License:        GPLv3
URL:            https://github.com/spideyz0r/kubesw
Source0:        %{url}/archive/refs/tags/v%{version}.tar.gz

BuildRequires:  golang >= %{go_version}
BuildRequires:  git

%description
kubesw is a cli tool to use switch namespaces and contexts per terminal

%global debug_package %{nil}

%prep
%autosetup -n %{name}-%{version}

%build
go build -v -o %{name} -ldflags=-linkmode=external

%check
go test

%install
install -Dpm 0755 %{name} %{buildroot}%{_bindir}/%{name}

%files
%{_bindir}/kubesw

%license LICENSE

%changelog
* Mon Jul 18 2023 spideyz0r <47341410+spideyz0r@users.noreply.github.com> 0.0.8-1
- Add keep history feature -- errata: removed 0.0.7

* Mon Jul 18 2023 spideyz0r <47341410+spideyz0r@users.noreply.github.com> 0.0.6-1
- Add new features

* Mon Jul 18 2023 spideyz0r <47341410+spideyz0r@users.noreply.github.com> 0.0.5-1
- Errata: fix modules and readme

* Mon Jul 17 2023 spideyz0r <47341410+spideyz0r@users.noreply.github.com> 0.0.4-1
- Code updates

* Mon Jul 17 2023 spideyz0r <47341410+spideyz0r@users.noreply.github.com> 0.0.3-1
- PS1 bugfix

* Mon Jul 17 2023 spideyz0r <47341410+spideyz0r@users.noreply.github.com> 0.0.2-1
- URL bugix

* Mon Jul 17 2023 spideyz0r <47341410+spideyz0r@users.noreply.github.com> 0.0.1-1
- Initial build

