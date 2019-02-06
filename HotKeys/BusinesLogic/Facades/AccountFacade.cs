using EnsureThat;
using HotKeysLibrary.Entities;
using HotKeysLibrary.Infrastructure;
using System;
using System.Collections.Generic;

namespace HotKeysLibrary.Facades
{
    public class AccountFacade
    {
        private readonly IRepository<User> _usersRepository;
        private readonly IRepository<Program> _programRepository;

        public AccountFacade(IRepository<User> usersRepository, IRepository<Program> programRepository)
        {
            _usersRepository = usersRepository;
            _programRepository = programRepository;
        }

        public Guid RegisterUser(string nickname)
        {
            Ensure.String.IsNotNullOrWhiteSpace(nickname);
            var id = Guid.NewGuid();
            var user = new User(id, nickname, new List<Program>());

            _usersRepository.Add(id, user);

            return id;

        }

    }
}
